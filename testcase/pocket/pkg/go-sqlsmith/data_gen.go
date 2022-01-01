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

package sqlsmith

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pingcap/tipocket/testcase/pocket/pkg/go-sqlsmith/types"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/go-sqlsmith/util"
)

// BatchData generate testing data by schema in given batch
// return SQLs with insert statement
func (s *SQLSmith) BatchData(total, batchSize int) ([]string, error) {
	if s.currDB == "" {
		return []string{}, errors.New("no selected database")
	}
	database, ok := s.Databases[s.currDB]
	if !ok {
		return []string{}, errors.New("selected database's schema not loaded")
	}

	var sqls []string
	for _, table := range database.Tables {
		var columns []*types.Column
		for _, column := range table.Columns {
			if column.Column == "id" {
				continue
			}
			columns = append(columns, column)
		}

		var lines [][]string
		count := 0
		for i := 0; i < total; i++ {
			var line []string
			for _, column := range columns {
				line = append(line, util.GenerateDataItemString(column.DataType))
	