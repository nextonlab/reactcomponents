module github.com/pingcap/tipocket/testcase/bank2

go 1.16

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9
	github.com/ngaut/log v0.0.0-20180314031856-b8e36e7ba5ac
	github.com/pingcap/tipocket v1.0.0
	github.com/pingcap/tipocket/logsearch v1.0.0
	github.com/rogpeppe/fastuuid v1.2.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/uber-go/atomic => go.uber.org/atomic v1.5.0

replace (
	k8s.io/api => k8s.io/api v0.17.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.0
	k8s.io/apimachinery => k