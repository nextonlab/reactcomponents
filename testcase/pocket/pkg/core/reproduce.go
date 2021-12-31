
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
	"bufio"
	"context"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
	"github.com/ngaut/log"

	"github.com/pingcap/tipocket/testcase/pocket/pkg/types"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/util"
)

var (
	abTestLogPattern     = regexp.MustCompile(`ab-test-[0-9]+\.log`)
	binlogTestLogPattern = regexp.MustCompile(`single-test-[0-9]+\.log`)
	todoSQLPattern       = regexp.MustCompile(`^\[([0-9]{4}\/[0-9]{2}\/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]{3} [+-][0-9]{2}:[0-9]{2})\] \[(TODO)\] Exec SQL (.*)$`)
	execIDPattern        = regexp.MustCompile(`^.*?(ab|single)-test-([0-9]+).log$`)
	timeLayout           = `2006/01/02 15:04:05.000 -07:00`
)

func (c *Core) reproduce(ctx context.Context) error {
	var (
		dir   = c.cfg.Options.Path
		table = ""
	)

	if dir == "" {
		log.Fatal("empty dir")
	} else if !util.DirExists(dir) {
		log.Fatal("invalid dir, not exist or not a dir")
	}
	return c.reproduceFromDir(dir, table)
}

func (c *Core) reproduceFromDir(dir, table string) error {
	var logFiles []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		match := false
		if c.cfg.Mode == "abtest" && abTestLogPattern.MatchString(f.Name()) {
			match = true
		} else if c.cfg.Mode == "binlog" && binlogTestLogPattern.MatchString(f.Name()) {
			match = true
		}
		if match {
			logFiles = append(logFiles, path.Join(dir, f.Name()))
		}
	}

	logs, err := c.readLogs(logFiles)
	if err != nil {
		return errors.Trace(err)
	}
