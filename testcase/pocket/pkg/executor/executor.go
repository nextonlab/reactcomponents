
// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package executor

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"sync"

	"github.com/juju/errors"
	"github.com/ngaut/log"

	"github.com/pingcap/tipocket/testcase/pocket/pkg/go-sqlsmith"

	"github.com/pingcap/tipocket/testcase/pocket/pkg/connection"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/generator/generator"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/logger"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/types"
)

var (
	dbnameRegex = regexp.MustCompile(`([a-z0-9A-Z_]+)$`)
)

// Executor define test executor
type Executor struct {
	sync.Mutex
	id          int
	dsn1        string
	dsn2        string
	dsn3        string
	conn1       *connection.Connection
	conn2       *connection.Connection
	conn3       *connection.Connection
	ss          generator.Generator
	dbname      string
	mode        string
	opt         *Option
	logger      *logger.Logger
	stmts       []*types.SQL
	stmtResults []*stmtResult
	// SQL channel
	SQLCh chan *types.SQL
	// Since we must ensure no other transactions commit or begin between the transaction start time points of abtest
	// when a transaction begins/commits/rollbacks, generator wait for it ready for both A/B side
	// This channel is for sending signal to generator when both A/B side's begin/commit/rollback are ready
	TxnReadyCh chan struct{}
	// ErrCh for waiting SQL execution finish
	// and pass execution error
	ErrCh chan error
	// OnlineTable record tables which are manipulated in transaction for avoiding online DDL
	OnlineTable []string
	TiFlash     bool
}

// New create Executor
func New(dsn string, opt *Option, tiFlash bool) (*Executor, error) {
	var connLogPath, executorLogPath string
	if opt.Log != "" {
		connLogPath = path.Join(opt.Log, fmt.Sprintf("single-conn%s.log", opt.LogSuffix))
		executorLogPath = path.Join(opt.Log, fmt.Sprintf("single-test%s.log", opt.LogSuffix))
	}

	conn, err := connection.New(dsn, &connection.Option{
		Name:          fmt.Sprintf("conn-%d", opt.ID),
		Log:           connLogPath,
		Mute:          opt.Mute,
		GeneralLog:    opt.GeneralLog,
		EnableTiFlash: tiFlash,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	l, err := logger.New(fmt.Sprintf("single-%d", opt.ID), executorLogPath, false)
	if err != nil {
		return nil, errors.Trace(err)
	}
	e := Executor{
		id:         opt.ID,
		dsn1:       dsn,
		conn1:      conn,
		mode:       "single",
		SQLCh:      make(chan *types.SQL, 1),
		TxnReadyCh: make(chan struct{}, 1),
		ErrCh:      make(chan error, 1),
		ss:         sqlsmith.New(),
		dbname:     dbnameRegex.FindString(dsn),
		opt:        opt,
		logger:     l,
		TiFlash:    tiFlash,
	}
	go e.Start()
	return &e, nil
}

// NewABTest create abtest Executor
func NewABTest(dsn1, dsn2 string, opt *Option, tiFlash bool) (*Executor, error) {
	var conn1LogPath, conn2LogPath, executorLogPath string
	if opt.Log != "" {
		conn1LogPath = path.Join(opt.Log, fmt.Sprintf("ab-conn1%s.log", opt.LogSuffix))
		conn2LogPath = path.Join(opt.Log, fmt.Sprintf("ab-conn2%s.log", opt.LogSuffix))
		executorLogPath = path.Join(opt.Log, fmt.Sprintf("ab-test%s.log", opt.LogSuffix))
	}

	conn1, err := connection.New(dsn1, &connection.Option{
		Name:          fmt.Sprintf("conn-1-%d", opt.ID),
		Log:           conn1LogPath,
		Mute:          opt.Mute,
		GeneralLog:    opt.GeneralLog,
		EnableTiFlash: tiFlash,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	conn2, err := connection.New(dsn2, &connection.Option{
		Name:       fmt.Sprintf("conn-2-%d", opt.ID),
		Log:        conn2LogPath,
		Mute:       opt.Mute,
		GeneralLog: opt.GeneralLog,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	l, err := logger.New(fmt.Sprintf("abtest-%d", opt.ID), executorLogPath, opt.Mute)
	if err != nil {
		return nil, errors.Trace(err)
	}
	e := Executor{
		id:         opt.ID,
		dsn1:       dsn1,
		dsn2:       dsn2,
		conn1:      conn1,
		conn2:      conn2,
		ss:         sqlsmith.New(),
		mode:       "abtest",
		SQLCh:      make(chan *types.SQL, 1),
		TxnReadyCh: make(chan struct{}, 1),
		ErrCh:      make(chan error, 1),
		dbname:     dbnameRegex.FindString(dsn1),
		opt:        opt,
		logger:     l,
		TiFlash:    tiFlash,
	}
	go e.Start()
	return &e, nil
}

// NewDMTest creates a DM test Executor.
// one `Start` for one executor.
func NewDMTest(dsn1, dsn2, dsn3 string, opt *Option, start bool) (*Executor, error) {
	var conn1LogPath, conn2LogPath, conn3LogPath, executorLogPath string
	if opt.Log != "" {
		conn1LogPath = path.Join(opt.Log, fmt.Sprintf("dm-conn1%s.log", opt.LogSuffix))
		conn2LogPath = path.Join(opt.Log, fmt.Sprintf("dm-conn2%s.log", opt.LogSuffix))
		conn3LogPath = path.Join(opt.Log, fmt.Sprintf("dm-conn3%s.log", opt.LogSuffix))
		executorLogPath = path.Join(opt.Log, fmt.Sprintf("dm-test%s.log", opt.LogSuffix))
	}

	conn1, err := connection.New(dsn1, &connection.Option{
		Name:       fmt.Sprintf("conn-1-%d", opt.ID),
		Log:        conn1LogPath,
		Mute:       opt.Mute,
		GeneralLog: opt.GeneralLog,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	log.Infof("new connection to %s", dsn1)
	conn2, err := connection.New(dsn2, &connection.Option{
		Name:       fmt.Sprintf("conn-2-%d", opt.ID),
		Log:        conn2LogPath,
		Mute:       opt.Mute,
		GeneralLog: opt.GeneralLog,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	log.Infof("new connection to %s", dsn2)
	conn3, err := connection.New(dsn3, &connection.Option{
		Name:       fmt.Sprintf("conn-3-%d", opt.ID),
		Log:        conn3LogPath,
		Mute:       opt.Mute,
		GeneralLog: opt.GeneralLog,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	log.Infof("new connection to %s", dsn3)

	l, err := logger.New(fmt.Sprintf("dmtest-%d", opt.ID), executorLogPath, opt.Mute)
	if err != nil {
		return nil, errors.Trace(err)
	}

	e := Executor{
		id:         opt.ID,
		dsn1:       dsn1,
		dsn2:       dsn2,
		dsn3:       dsn3,
		conn1:      conn1,
		conn2:      conn2,
		conn3:      conn3,
		ss:         sqlsmith.New(),
		mode:       "dm",
		SQLCh:      make(chan *types.SQL), // no buffer for the chan to ensure `execMutex` can really lock exec.
		TxnReadyCh: make(chan struct{}, 1),
		ErrCh:      make(chan error, 1),
		dbname:     dbnameRegex.FindString(dsn1),
		opt:        opt,
		logger:     l,
	}
	if start {
		go e.Start()
	}
	return &e, nil
}

// func (e *Executor) init() error {
// 	switch e.mode {
// 	case "single":
// 		return errors.Trace(e.singleTestReloadSchema())
// 	case "abtest":
// 		return errors.Trace(e.abTestReloadSchema())
// 	}
// 	return errors.New("not support mode")
// }

// Start start test
func (e *Executor) Start() {
	// if err := e.init(); err != nil {
	// 	log.Fatalf("init failed %v\n", errors.ErrorStack(err))
	// }
	// if e.opt.Reproduce != "" {
	// 	go e.reproduce()
	// } else {
	// 	go e.smithGenerate()
	// }
	switch e.mode {
	case "single":
		e.singleTest()
	case "abtest":
		e.abTest()
	case "dm":
		e.dmTest()
	}
}

// Stop exit process
func (e *Executor) Stop(msg string) {
	log.Infof("[STOP] message: %s\n", msg)
	os.Exit(0)
}

// GetID return executor id
func (e *Executor) GetID() int {
	return e.id
}