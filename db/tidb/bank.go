
package tidb

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/anishathalye/porcupine"

	pchecker "github.com/pingcap/tipocket/pkg/check/porcupine"
	"github.com/pingcap/tipocket/pkg/cluster"
	"github.com/pingcap/tipocket/pkg/core"
	"github.com/pingcap/tipocket/pkg/history"

	// use mysql
	_ "github.com/go-sql-driver/mysql"
)

const (
	accountNum  = 5
	initBalance = int64(1000)
)

type bankClient struct {
	db         *sql.DB
	r          *rand.Rand
	accountNum int
}

func (c *bankClient) SetUp(ctx context.Context, _ []cluster.Node, clientNodes []cluster.ClientNode, idx int) error {
	c.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	node := clientNodes[idx]
	db, err := sql.Open("mysql", fmt.Sprintf("root@tcp(%s:%d)/test", node.IP, node.Port))
	if err != nil {
		return err
	}
	c.db = db

	db.SetMaxIdleConns(1 + c.accountNum)

	// Do SetUp in the first node
	if idx != 0 {
		return nil
	}

	log.Printf("begin to create table accounts on node %+v", node)
	if _, err = db.ExecContext(ctx, `drop table if exists accounts`); err != nil {
		return err
	}

	sql := `create table if not exists accounts
			(id     int not null primary key,
			balance bigint not null)`

	if _, err = db.ExecContext(ctx, sql); err != nil {
		return err
	}

	for i := 0; i < c.accountNum; i++ {
		if _, err = db.ExecContext(ctx, "insert into accounts values (?, ?)", i, initBalance); err != nil {
			return err
		}
	}

	return nil
}

func (c *bankClient) TearDown(ctx context.Context, nodes []cluster.ClientNode, idx int) error {
	if idx != 0 {
		return nil
	}
	sql := `drop table if exists accounts`
	if _, err := c.db.Exec(sql); err != nil {
		return err
	}
	return c.db.Close()
}

func (c *bankClient) invokeRead(ctx context.Context, r bankRequest) bankResponse {
	txn, err := c.db.Begin()

	if err != nil {
		return bankResponse{Unknown: true}
	}
	defer txn.Rollback()

	var tso uint64
	if err = txn.QueryRow("select @@tidb_current_ts").Scan(&tso); err != nil {
		return bankResponse{Unknown: true}
	}

	rows, err := txn.QueryContext(ctx, "select balance from accounts")
	if err != nil {
		return bankResponse{Unknown: true, Error: err.Error()}
	}
	defer rows.Close()

	balances := make([]int64, 0, c.accountNum)
	for rows.Next() {
		var v int64
		if err = rows.Scan(&v); err != nil {
			return bankResponse{Unknown: true, Error: err.Error()}
		}
		balances = append(balances, v)
	}

	return bankResponse{Balances: balances, Tso: tso}
}

func (c *bankClient) Invoke(ctx context.Context, node cluster.ClientNode, r interface{}) core.UnknownResponse {
	arg := r.(bankRequest)
	if arg.Op == 0 {
		return c.invokeRead(ctx, arg)
	}

	txn, err := c.db.Begin()

	if err != nil {
		return bankResponse{Ok: false}
	}
	defer txn.Rollback()

	var (
		fromBalance int64
		toBalance   int64
		tso         uint64
	)

	if err = txn.QueryRow("select @@tidb_current_ts").Scan(&tso); err != nil {
		return bankResponse{Ok: false}
	}

	if err = txn.QueryRowContext(ctx, "select balance from accounts where id = ? for update", arg.From).Scan(&fromBalance); err != nil {
		return bankResponse{Ok: false}
	}

	if err = txn.QueryRowContext(ctx, "select balance from accounts where id = ? for update", arg.To).Scan(&toBalance); err != nil {
		return bankResponse{Ok: false}
	}

	if fromBalance < arg.Amount {
		return bankResponse{Ok: false}
	}

	if _, err = txn.ExecContext(ctx, "update accounts set balance = balance - ? where id = ?", arg.Amount, arg.From); err != nil {
		return bankResponse{Ok: false}
	}

	if _, err = txn.ExecContext(ctx, "update accounts set balance = balance + ? where id = ?", arg.Amount, arg.To); err != nil {
		return bankResponse{Ok: false}
	}

	if err = txn.Commit(); err != nil {
		return bankResponse{Unknown: true, Tso: tso, FromBalance: fromBalance, ToBalance: toBalance, Error: err.Error()}
	}

	return bankResponse{Ok: true, Tso: tso, FromBalance: fromBalance, ToBalance: toBalance}
}

func (c *bankClient) NextRequest() interface{} {
	r := bankRequest{
		Op: c.r.Int() % 2,
	}
	if r.Op == 0 {
		return r
	}

	r.From = c.r.Intn(c.accountNum)

	r.To = c.r.Intn(c.accountNum)
	if r.From == r.To {
		r.To = (r.To + 1) % c.accountNum
	}

	r.Amount = 5
	return r
}

func (c *bankClient) DumpState(ctx context.Context) (interface{}, error) {
	txn, err := c.db.Begin()

	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	rows, err := txn.QueryContext(ctx, "select balance from accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	balances := make([]int64, 0, c.accountNum)
	for rows.Next() {
		var v int64
		if err = rows.Scan(&v); err != nil {
			return nil, err
		}
		balances = append(balances, v)
	}
	return balances, nil
}

// BankClientCreator creates a bank test client for tidb.
type BankClientCreator struct {
}

// Create creates a client.
func (BankClientCreator) Create(node cluster.ClientNode) core.Client {
	return &bankClient{
		accountNum: accountNum,
		r:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type bankRequest struct {
	// 0: read
	// 1: transfer
	Op     int
	From   int
	To     int
	Amount int64
}

type bankResponse struct {
	// Transaction start timestamp
	Tso uint64
	// read result
	Balances []int64
	// transfer ok or not
	Ok bool
	// FromBalance is the previous from balance before transfer
	FromBalance int64
	// ToBalance is the previous to balance before transfer
	ToBalance int64
	// read/transfer unknown
	Unknown bool
	// record the error if Unknown
	Error string `json:",omitempty"`
}

var _ core.UnknownResponse = (*bankResponse)(nil)

// IsUnknown implements UnknownResponse interface
func (r bankResponse) IsUnknown() bool {
	return r.Unknown
}

func balancesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type bank struct {
	accountNum    int
	perparedState *[]int64
}

func (b *bank) Prepare(state interface{}) {
	s := state.([]int64)
	b.perparedState = &s
}

func (b *bank) Init() interface{} {
	if b.perparedState != nil {
		return *b.perparedState
	}

	// Or make a brand new state.
	v := make([]int64, b.accountNum)
	for i := 0; i < b.accountNum; i++ {
		v[i] = initBalance
	}
	return v
}

func (*bank) Step(state interface{}, input interface{}, output interface{}) (bool, interface{}) {
	st := state.([]int64)
	inp := input.(bankRequest)
	out := output.(bankResponse)

	if inp.Op == 0 {
		// read
		ok := out.Unknown || balancesEqual(st, out.Balances)
		return ok, state
	}

	// for transfer
	if !out.Ok && !out.Unknown {
		return true, state
	}

	newSt := append([]int64{}, st...)
	newSt[inp.From] -= inp.Amount
	newSt[inp.To] += inp.Amount
	return out.Ok || out.Unknown, newSt
}

func (*bank) Equal(state1, state2 interface{}) bool {
	st1 := state1.([]int64)
	st2 := state2.([]int64)
	return balancesEqual(st1, st2)
}

func (*bank) Name() string {
	return "tidb_bank"
}

// BankModel is the model of bank in TiDB
func BankModel() core.Model {
	return &bank{
		accountNum: accountNum,
	}
}

type bankParser struct{}

// OnRequest impls history.RecordParser.OnRequest
func (p bankParser) OnRequest(data json.RawMessage) (interface{}, error) {
	r := bankRequest{}
	err := json.Unmarshal(data, &r)
	return r, err
}

// OnResponse impls history.RecordParser.OnRequest
func (p bankParser) OnResponse(data json.RawMessage) (interface{}, error) {
	r := bankResponse{}
	err := json.Unmarshal(data, &r)
	return r, err
}

// OnNoopResponse impls history.RecordParser.OnRequest
func (p bankParser) OnNoopResponse() interface{} {
	return bankResponse{Unknown: true}
}

func (p bankParser) OnState(data json.RawMessage) (interface{}, error) {
	var state []int64
	err := json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// BankParser parses a history of bank operations.
func BankParser() history.RecordParser {
	return bankParser{}
}

type tsoEvent struct {
	Tso uint64
	Op  int
	// For transfer
	From        int
	To          int
	FromBalance int64
	ToBalance   int64
	Amount      int64
	// For read
	Balances []int64

	Unknown bool
}

func (e *tsoEvent) String() string {
	if e.Op == 0 {
		return fmt.Sprintf("%d, read %v, unknown %v", e.Tso, e.Balances, e.Unknown)
	}

	return fmt.Sprintf("%d, transafer %d %d(%d) -> %d(%d), unknown %v", e.Tso, e.Amount, e.From, e.FromBalance, e.To, e.ToBalance, e.Unknown)
}

// GetBalances gets the two balances of account before and after the transfer.
func (e *tsoEvent) GetBalances(index int) (int64, int64) {
	if index == e.From {
		return e.FromBalance, e.FromBalance - e.Amount
	}

	return e.ToBalance, e.ToBalance + e.Amount
}

type tsoEvents []*tsoEvent

func (s tsoEvents) Len() int           { return len(s) }
func (s tsoEvents) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s tsoEvents) Less(i, j int) bool { return s[i].Tso < s[j].Tso }

// TODO: remove porcupine dependence.
func generateTsoEvents(events []porcupine.Event) tsoEvents {
	tEvents := make(tsoEvents, 0, len(events))

	mapEvents := make(map[int]porcupine.Event, len(events))
	for _, event := range events {
		if event.Kind == porcupine.CallEvent {
			mapEvents[event.Id] = event
			continue
		}

		// Handle Return Event
		// Find the corresponding Call Event
		callEvent, ok := mapEvents[event.Id]
		if !ok {
			continue
		}
		delete(mapEvents, event.Id)

		request := callEvent.Value.(bankRequest)
		response := event.Value.(bankResponse)

		if response.Tso == 0 {
			// We don't care operation which has no TSO.
			continue
		}

		tEvent := tsoEvent{
			Tso:     response.Tso,
			Op:      request.Op,
			Unknown: response.Unknown,
		}
		if request.Op == 0 {
			tEvent.Balances = response.Balances
		} else {
			tEvent.From = request.From
			tEvent.To = request.To
			tEvent.Amount = request.Amount
			tEvent.FromBalance = response.FromBalance
			tEvent.ToBalance = response.ToBalance
		}

		tEvents = append(tEvents, &tEvent)
	}
	sort.Sort(tEvents)
	return tEvents
}

// mergeTransferEvents checks whether e can be merged into the events.
// We may meet following cases for one account:
// Assume last event starts at T1, the checking event starts at T2.
// 1:
// 	T1: [1000] -> [900], Unknown
//	T2: [900] -> [800], Unknown?
// Here T2 reads 900, so we can ensure T1 is successful no matter T1 is unknown or not.
// We can set T1 to OK. After T1 is set to OK, we must check T1 to its previous events.
// 2:
//	T1: [1000] -> [900], OK
//	T2: [1000] -> [800], Unknown
// Here T1 is successful, but T2 is unknown, it is fine now.
// 3:
// 	T1: [1000] -> [900], Ok
//	T2: [1000] -> [800], Ok
// Invalid, because we use SSI here, even T2 can read 1000, it can't change it because
// it must conflict with T1.
// 4:
// 	T1: [1000] -> [900], Unknown?
//	T2: [800] -> [700], Unknown?
// Invalid, T2 reads a stale value.
func mergeTransferEvents(index int, events tsoEvents, e *tsoEvent) (tsoEvents, error) {
	curBalance, _ := e.GetBalances(index)

	if !checkBalance(index, events, curBalance) {
		return nil, fmt.Errorf("%d %v invalid event %s", index, events, e)
	}

	events = append(events, e)

	// Get the last successful event e2
	lastIdx, err := checkTransferEvents(index, events)
	if err != nil {
		return nil, err
	}

	// clear all events before the successful event
	return events[lastIdx:], nil
}

// For all the successful transfer events, we must form a transfer chain like
// T1 [1000] -> [900]
// T2 [900] -> [800]
// T3 [800] -> [700]
// The function will return the last successful event index, if no found, return 0
func checkTransferEvents(index int, events tsoEvents) (int, error) {
	var (
		lastEvent *tsoEvent
		lastIndex int
	)
	for i, e := range events {
		if e.Unknown {
			continue
		}

		if lastEvent != nil {
			_, next := lastEvent.GetBalances(index)
			cur, _ := e.GetBalances(index)
			if next != cur {
				return 0, fmt.Errorf("invalid events from %s to %s", lastEvent, e)
			}
		}

		lastIndex = i
		lastEvent = e
	}

	return lastIndex, nil
}

func checkBalance(index int, events tsoEvents, curBalance int64) bool {
	if len(events) == 0 {
		return curBalance == initBalance
	}

	for i := len(events) - 1; i >= 0; i-- {
		lastEvent := events[i]
		cur, next := lastEvent.GetBalances(index)
		if next == curBalance {
			// We read the next balance of the last event, which means the last transfer is
			// successful
			lastEvent.Unknown = false
			return true
		}

		if cur == curBalance {
			// Oh, we read the same balance with the last event
			return true
		}
	}

	return false
}

// verifyReadEvent verifies the read event.
func verifyReadEvent(possibleEvents []tsoEvents, e *tsoEvent) bool {
	if e.Unknown {
		return true
	}

	sum := int64(0)
	for i, balance := range e.Balances {
		sum += balance

		if !checkBalance(i, possibleEvents[i], balance) {
			log.Printf("invalid event %s, balance mismatch", e)
			return false
		}
	}

	if sum != int64(len(e.Balances))*initBalance {
		log.Printf("invalid event %s, sum corruption", e)
		return false
	}

	return true
}

func verifyTsoEvents(events tsoEvents) bool {
	possibleEvents := make([]tsoEvents, accountNum)

	var err error
	for _, event := range events {
		if event.Op == 0 {
			if !verifyReadEvent(possibleEvents, event) {
				return false
			}
		}

		if event.Op == 1 {
			from := event.From
			possibleEvents[from], err = mergeTransferEvents(from, possibleEvents[from], event)
			if err != nil {
				log.Print(err.Error())
				return false
			}

			to := event.To
			possibleEvents[to], err = mergeTransferEvents(to, possibleEvents[to], event)
			if err != nil {
				log.Print(err.Error())
				return false
			}
		}
	}

	return true
}

// bankTsoChecker uses a direct way because we know every timestamp of the transaction.
// So we can order all transactions with timetamp and replay them.
type bankTsoChecker struct {
}

// Check checks the bank history.
func (bankTsoChecker) Check(_ core.Model, ops []core.Operation) (bool, error) {
	events, err := pchecker.ConvertOperationsToEvents(ops)
	if err != nil {
		return false, err
	}
	tEvents := generateTsoEvents(events)
	return verifyTsoEvents(tEvents), nil
}

// Name returns the name of the verifier.
func (bankTsoChecker) Name() string {
	return "tidb_bank_tso_checker"
}

// BankTsoChecker checks the bank history with the help of tso.
func BankTsoChecker() core.Checker {
	return bankTsoChecker{}
}