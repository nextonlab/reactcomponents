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

package binlog

import (
	"html/template"

	"github.com/pingcap/tipocket/pkg/test-infra/util"
)

// copy from https://github.com/pingcap/tidb-binlog/blame/e28b75cac81bea82c2a89ad024d1a37bf3c9bee9/cmd/drainer/drainer.toml#L43
var drainerConfigTpl = template.Must(template.New("drainer-config-script").Parse(`# drainer Configuration.

# addr (i.e. 'host:port') to listen on for drainer con