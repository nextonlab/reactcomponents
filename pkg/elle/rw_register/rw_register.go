package rwregister

import (
	"fmt"
	"log"
	"sort"

	"github.com/pingcap/tipocket/pkg/elle/core"
	"github.com/pingcap/tipocket/pkg/elle/txn"
)

// GraphOption ...
type GraphOption struct {
	LinearizableKeys bool //Uses realtime order
	SequentialKeys   bool // Uses process order
	WfrKeys          bool // Assumes writes follow reads in a txn
}

// GCaseTp type aliases []core.Anomaly
type GCaseTp []core.Anomaly

// InternalConflict records a internal conflict
type InternalConflict struct {
	Op       core.Op
	Mop      core.Mop
	Expected core.Mop
}

// IAnomaly ...
func (i InternalConflict) IAnomaly() {}

// String ...
func (i InternalConflict) String() string {
	return 