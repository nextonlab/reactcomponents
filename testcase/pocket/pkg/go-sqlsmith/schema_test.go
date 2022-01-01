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
	"testing"

	"github.com/stretchr/testify/assert"
)

// test schema
var (
	schema = [][5]string{
		{"community", "comments", "BASE TABLE", "id", "int(11)"},
		{"community", "comments", "BASE TABLE", "owner", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "repo", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "comment_id", "int(11)"},
		{"community", "comments", "BASE TABLE", "comment_type", "varchar(128)"},
		{"community", "comments", "BASE TABLE", "pull_number", "int(11)"},
		{"community", "comments", "BASE TABLE", "body", "text"},
		{"community", "comments", "BASE TABLE", "user", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "url", "varchar(1023)"},
		{"community", "comments", "BASE TABLE", "association", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "relation", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "created_at", "timestamp"},
		{"community", "comments", "BASE TABLE", "updated_at", "timestamp"},
		{"community", "picks", "BASE TABLE", "id", "int(11)"},
		{"community", "picks", "BASE TABLE", "season", "int(11)"},
		{"community", "picks", "BASE TABLE", "task_id", "int(11)"},
		{"community", "picks", "BASE TABLE", "teamID", "int(11)"},
		{"community", "picks", "BASE TABLE", "user", "varchar(255)"},
		{"community", "picks", "BASE TABLE", "pull_number", "int(11)"},
		{"community", "picks", "BASE TABLE", "status", "varchar(128)"},
		{"community", "picks", "BASE TABLE", "created_at", "timestamp"},
		{"community", "picks", "BASE TABLE", "updated_at", "timestamp"},
		{