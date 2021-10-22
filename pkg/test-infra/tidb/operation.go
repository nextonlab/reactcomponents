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

package tidb

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/ngaut/log"
	"github.com/pingcap/errors"
	"golang.org/x/sync/errgroup"

	"github.com/pingcap/tipocket/pkg/cluster"
	"github.com/pingcap/tipocket/pkg/test-infra/fixture"
	"github.com/pingcap/tipocket/pkg/test-infra/tests"
	"github.com/pingcap/tipocket/pkg/test-infra/util"
	"github.com/pingcap/tipocket/pkg/tidb-operator/apis/pingcap/v1alpha1"
	"github.com/pingcap/tipocket/pkg/tidb-operator/label"
)

const (
	tikvDir     = "/var/lib/tikv"
	tikvDataDir = "/var/lib/tikv/data"
	pdDir       = "/var/lib/pd"
	pdDataDir   = "/var/lib/pd/data"
	// used for tikv data encryption
	tikvEncryptionMasterKey = "c7fd825f4ec91c07067553896cb1b4ad9e32e9175e7750aa39cc1771fc8eb589"
	plaintextProtocolHeader = "plaintext://"
)

// Ops knows how to operate TiDB
type Ops struct {
	cli    client.Client
	tc     *Recommendation
	config fixture.TiDBClusterConfig
	ns