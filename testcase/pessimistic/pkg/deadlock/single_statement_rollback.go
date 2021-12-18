package deadlock

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ngaut/log"
	"github.com/pingcap/errors"

	"github.com/pingcap/tipocket/util"
)

type singleStatementRollbackCase struct {
	dsn      string
	tableNum int
	interval time.Duration
	db       *sql.DB
	query    string
	chs      []chan struct{}
}

func newSingleStatementRollbackCase(dsn string, tableNum int, interval time.Duration) *singleStatementRollbackCase {
	return &singleStatementRollbackCase{
		dsn:      dsn,
		tableNum: tableNum,
		interval: interval,
	}
}

func (s *singleStatementRollbackCase) initialize(ctx context.Context) error {
	log.Info("[single statement rollback] Initialize database")
	if err := s.check(); err != nil {
		return err
	}
	if err := s.openDB(); err != nil {
		return err
	}
	if err := s.createTables(); err != nil {
		return err
	}
	return nil
}

func (s *singleStatementRollbackCase) check() error {
	if s.tableNum <= 1 {
		return errors.New("tableNum should be greater than 1")
	}
	// TODO(yeya24): add that check once we can get tikv config.
	// This case should check 'split-region-on-table'.
	// However we can't get tikv config through http api in release-3.0.
	return nil
}

func (s *singleStatementRollbackCase) openDB() error {
	// Connect to DB
	log.Infof("[deadlock] DSN: %s", s.dsn)
	var err error
	s.db, err = util.OpenDB(s.dsn, 2)
	if err != nil {
		return errors.Wrap(err, "Open db failed")
	}
	return nil
}

func (s *singleStatementRollbackCase) createTa