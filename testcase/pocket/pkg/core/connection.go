
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

package core

import (
	"fmt"
	"strings"

	"github.com/juju/errors"

	"github.com/pingcap/tipocket/testcase/pocket/pkg/executor"
)

func (c *Core) generateExecutorOption(id int) *executor.Option {
	var suffix string
	if id > 0 {
		suffix = fmt.Sprintf("-%d", id)
	}
	opt := executor.Option{
		ID:         id,
		Log:        c.cfg.Options.Path,
		LogSuffix:  suffix,
		Stable:     c.cfg.Options.Stable,
		Mute:       c.cfg.Options.Reproduce,
		OnlineDDL:  c.cfg.Options.OnlineDDL,
		GeneralLog: c.cfg.Options.GeneralLog,
		Hint:       c.cfg.Options.EnableHint,
	}
	return &opt
}

func (c *Core) initConnectionWithoutSchema(id int) (*executor.Executor, error) {
	var (
		e       *executor.Executor
		err     error
		tiFlash bool
	)

	// current supported tiflash mode includes tiflash, tiflash-abtest, tiflash-binlog
	if strings.HasPrefix(c.cfg.Mode, "tiflash") {
		tiFlash = true
	}

	switch c.cfg.Mode {
	case "single", "tiflash":
		e, err = executor.New(removeDSNSchema(c.cfg.DSN1), c.generateExecutorOption(id), tiFlash)
		if err != nil {
			return nil, errors.Trace(err)
		}
	case "abtest", "binlog", "tiflash-abtest", "tiflash-binlog":
		e, err = executor.NewABTest(removeDSNSchema(c.cfg.DSN1),
			removeDSNSchema(c.cfg.DSN2),
			c.generateExecutorOption(id), tiFlash)
		if err != nil {
			return nil, errors.Trace(err)