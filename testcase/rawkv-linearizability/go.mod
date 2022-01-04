
module github.com/pingcap/tipocket/testcase/rawkv-linearizability

go 1.16

require (
	github.com/gengliqi/persistent_treap v0.0.0-20200403155416-2b2a1532211c
	github.com/pingcap/kvproto v0.0.0-20200324130106-b8bc94dd8a36 // indirect
	github.com/pingcap/pd v2.1.17+incompatible
	github.com/pingcap/tipocket v1.0.0
	github.com/tikv/client-go v0.0.0-20200110101306-a3ebdb020c83
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/uber-go/atomic => go.uber.org/atomic v1.5.0

replace (
	k8s.io/api => k8s.io/api v0.17.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.0
	k8s.io/apiserver => k8s.io/apiserver v0.17.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.0
	k8s.io/client-go => k8s.io/client-go v0.17.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.0
	k8s.io/code-generator => k8s.io/code-generator v0.17.0
	k8s.io/component-base => k8s.io/component-base v0.17.0
	k8s.io/cri-api => k8s.io/cri-api v0.17.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.0
	k8s.io/kubectl => k8s.io/kubectl v0.17.0
	k8s.io/kubelet => k8s.io/kubelet v0.17.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.0
	k8s.io/metrics => k8s.io/metrics v0.17.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible

replace golang.org/x/net v0.0.0-20190813000000-74dc4d7220e7 => golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7

replace github.com/pingcap/tipocket => ../../.