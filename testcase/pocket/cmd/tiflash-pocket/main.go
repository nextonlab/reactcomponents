
// Copyright 2020 PingCAP, Inc.
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

package main

import (
	"context"
	"flag"

	"github.com/pingcap/tipocket/cmd/util"
	"github.com/pingcap/tipocket/pkg/cluster"
	"github.com/pingcap/tipocket/pkg/control"
	test_infra "github.com/pingcap/tipocket/pkg/test-infra"
	"github.com/pingcap/tipocket/pkg/test-infra/fixture"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/config"
	"github.com/pingcap/tipocket/testcase/pocket/pkg/creator"
)

var (
	configPath = flag.String("config", "", "config file path")
)

func main() {
	flag.Parse()
	cfg := control.Config{
		Mode:        control.ModeStandard,
		ClientCount: 1,
		RunTime:     fixture.Context.RunTime,
		RunRound:    1,
	}
	pocketConfig := config.Init()
	pocketConfig.Options.Serialize = false
	pocketConfig.Options.Path = "tiflash.log"
	pocketConfig.Options.EnableHint = fixture.Context.EnableHint
	suit := util.Suit{
		Config:   &cfg,
		Provider: cluster.NewDefaultClusterProvider(),
		ClientCreator: creator.PocketCreator{
			Config: creator.Config{
				ConfigPath: *configPath,
				Mode:       "tiflash",
				Config:     pocketConfig,
			},
		},
		NemesisGens:      util.ParseNemesisGenerators(fixture.Context.Nemesis),
		ClientRequestGen: util.OnClientLoop,
		ClusterDefs: test_infra.NewTiFlashCluster(fixture.Context.Namespace, fixture.Context.ClusterName,
			fixture.Context.TiDBClusterConfig),
	}
	suit.Run(context.Background())
}