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

package util

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ngaut/log"

	"github.com/pingcap/tipocket/pkg/cluster"
	"github.com/pingcap/tipocket/pkg/control"
	"github.com/pingcap/tipocket/pkg/core"
	"github.com/pingcap/tipocket/pkg/history"
	"github.com/pingcap/tipocket/pkg/logs"
	"github.com/pingcap/tipocket/pkg/nemesis"
	"github.com/pingcap/tipocket/pkg/test-infra/fixture"
	"github.com/pingcap/tipocket/pkg/verify"
)

// Suit is a basic chaos testing suit with configurations to run chaos.
type Suit struct {
	*control.Config
	// Provider deploy the SUT cluster
	cluster.Provider
	// ClientCreator creates client
	core.ClientCreator
	// NemesisGens saves NemesisGenerator
	NemesisGens []core.NemesisGenerator
	// ClientRequestGen
	ClientRequestGen ClientLoopFunc
	// perform service quality checking
	VerifySuit verify.Suit
	// cluster definition
	ClusterDefs cluster.Cluster
	// Plugins
	Plugins []control.Plugin
	// LogsSearch client
	LogsClient logs.SearchLogClient
}

// Run runs the suit.
func (suit *Suit) Run(ctx context.Context) {
	var (
		err         error
		clusterSpec = 