
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

package v1alpha1

// Port from TiFlash configurations till 2020/04/02

// TiFlashConfig is the configuration of TiFlash.
// +k8s:openapi-gen=true
type TiFlashConfig struct {
	// commonConfig is the Configuration of TiFlash process
	// +optional
	CommonConfig *CommonConfig `json:"config,omitempty"`

	// proxyConfig is the Configuration of proxy process
	// +optional
	ProxyConfig *ProxyConfig `json:"proxy,omitempty"`
}

// FlashServerConfig is the configuration of Proxy server.
// +k8s:openapi-gen=true
type FlashServerConfig struct {
	// Default to {clusterName}-tiflash-POD_NUM.{clusterName}-tiflash-peer.{namespace}:3930
	// +optional
	EngineAddr *string `json:"engine-addr,omitempty" toml:"engine-addr,omitempty"`
	// Default to 0.0.0.0:20292
	// +optional
	StatusAddr *string `json:"status-addr,omitempty" toml:"status-addr,omitempty"`
	// Default to {clusterName}-tiflash-POD_NUM.{clusterName}-tiflash-peer.{namespace}:20292
	// +optional
	AdvertiseStatusAddr *string `json:"advertise-status-addr,omitempty" toml:"advertise-status-addr,omitempty"`
	TiKVServerConfig    `json:",inline"`
}

// ProxyConfig is the configuration of TiFlash proxy process.
// All the configurations are same with those of TiKV except adding `engine-addr` in the TiKVServerConfig
// +k8s:openapi-gen=true
type ProxyConfig struct {
	// Optional: Defaults to info
	// +optional
	LogLevel *string `json:"log-level,omitempty" toml:"log-level,omitempty"`
	// +optional
	LogFile *string `json:"log-file,omitempty" toml:"log-file,omitempty"`
	// Optional: Defaults to 24h
	// +optional
	LogRotationTimespan *string `json:"log-rotation-timespan,omitempty" toml:"log-rotation-timespan,omitempty"`
	// +optional
	PanicWhenUnexpectedKeyOrData *bool `json:"panic-when-unexpected-key-or-data,omitempty" toml:"panic-when-unexpected-key-or-data,omitempty"`
	// +optional
	Server *FlashServerConfig `json:"server,omitempty" toml:"server,omitempty"`
	// +optional
	Storage *TiKVStorageConfig `json:"storage,omitempty" toml:"storage,omitempty"`
	// +optional
	Raftstore *TiKVRaftstoreConfig `json:"raftstore,omitempty" toml:"raftstore,omitempty"`
	// +optional
	Rocksdb *TiKVDbConfig `json:"rocksdb,omitempty" toml:"rocksdb,omitempty"`
	// +optional
	Coprocessor *TiKVCoprocessorConfig `json:"coprocessor,omitempty" toml:"coprocessor,omitempty"`
	// +optional
	ReadPool *TiKVReadPoolConfig `json:"readpool,omitempty" toml:"readpool,omitempty"`
	// +optional
	RaftDB *TiKVRaftDBConfig `json:"raftdb,omitempty" toml:"raftdb,omitempty"`
	// +optional
	Import *TiKVImportConfig `json:"import,omitempty" toml:"import,omitempty"`
	// +optional
	GC *TiKVGCConfig `json:"gc,omitempty" toml:"gc,omitempty"`
	// +optional
	PD *TiKVPDConfig `json:"pd,omitempty" toml:"pd,omitempty"`
	// +optional
	Security *TiKVSecurityConfig `json:"security,omitempty" toml:"security,omitempty"`
}

// CommonConfig is the configuration of TiFlash process.
// +k8s:openapi-gen=true
type CommonConfig struct {
	// Optional: Defaults to "/data0/tmp"
	// +optional
	TmpPath *string `json:"tmp_path,omitempty" toml:"tmp_path,omitempty"`

	// Optional: Defaults to "TiFlash"
	// +optional
	// +k8s:openapi-gen=false
	DisplayName *string `json:"display_name,omitempty" toml:"display_name,omitempty"`

	// Optional: Defaults to "default"
	// +optional
	// +k8s:openapi-gen=false
	DefaultProfile *string `json:"default_profile,omitempty" toml:"default_profile,omitempty"`

	// Optional: Defaults to "/data0/db"
	// +optional
	// +k8s:openapi-gen=false
	FlashDataPath *string `json:"path,omitempty" toml:"path,omitempty"`

	// Optional: Defaults to false
	// +optional
	PathRealtimeMode *bool `json:"path_realtime_mode,omitempty" toml:"path_realtime_mode,omitempty"`

	// Optional: Defaults to 5368709120
	// +optional
	MarkCacheSize *int64 `json:"mark_cache_size,omitempty" toml:"mark_cache_size,omitempty"`

	// Optional: Defaults to 5368709120
	// +optional
	MinmaxIndexCacheSize *int64 `json:"minmax_index_cache_size,omitempty" toml:"minmax_index_cache_size,omitempty"`

	// Optional: Defaults to "0.0.0.0"
	// +optional
	// +k8s:openapi-gen=false
	ListenHost *string `json:"listen_host,omitempty" toml:"listen_host,omitempty"`

	// Optional: Defaults to 9000
	// +optional
	// +k8s:openapi-gen=false
	TCPPort *int32 `json:"tcp_port,omitempty" toml:"tcp_port,omitempty"`
	// Optional: Defaults to 8123
	// +optional
	// +k8s:openapi-gen=false
	HTTPPort *int32 `json:"http_port,omitempty" toml:"http_port,omitempty"`
	// Optional: Defaults to 9000
	// +optional
	// +k8s:openapi-gen=false
	TCPPortSecure *int32 `json:"tcp_port_secure,omitempty" toml:"tcp_port_secure,omitempty"`
	// Optional: Defaults to 8123
	// +optional
	// +k8s:openapi-gen=false
	HTTPSPort *int32 `json:"https_port,omitempty" toml:"https_port,omitempty"`
	// Optional: Defaults to 9009
	// +optional
	// +k8s:openapi-gen=false
	InternalServerHTTPPort *int32 `json:"interserver_http_port,omitempty" toml:"interserver_http_port,omitempty"`
	// +optional
	Flash *Flash `json:"flash,omitempty" toml:"flash,omitempty"`
	// +optional
	FlashLogger *FlashLogger `json:"logger,omitempty" toml:"logger,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	FlashApplication *FlashApplication `json:"application,omitempty" toml:"application,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	FlashRaft *FlashRaft `json:"raft,omitempty" toml:"raft,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	FlashStatus *FlashStatus `json:"status,omitempty" toml:"status,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	FlashQuota *FlashQuota `json:"quotas,omitempty" toml:"quotas,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	FlashUser *FlashUser `json:"users,omitempty" toml:"users,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	FlashProfile *FlashProfile `json:"profiles,omitempty" toml:"profiles,omitempty"`
	// +optional
	// +k8s:openapi-gen=false
	Security *FlashSecurity `json:"security,omitempty" toml:"security,omitempty"`
}

// FlashProfile is the configuration of [profiles] section.
// +k8s:openapi-gen=false
type FlashProfile struct {
	// +optional
	Readonly *Profile `json:"readonly,omitempty" toml:"readonly,omitempty"`
	// +optional
	Default *Profile `json:"default,omitempty" toml:"default,omitempty"`
}

// Profile is the configuration profiles.
// +k8s:openapi-gen=false
type Profile struct {
	// +optional
	Readonly *int32 `json:"readonly,omitempty" toml:"readonly,omitempty"`
	// +optional
	MaxMemoryUsage *int64 `json:"max_memory_usage,omitempty" toml:"max_memory_usage,omitempty"`
	// +optional
	UseUncompressedCache *int32 `json:"use_uncompressed_cache,omitempty" toml:"use_uncompressed_cache,omitempty"`
	// +optional
	LoadBalancing *string `json:"load_balancing,omitempty" toml:"load_balancing,omitempty"`
}

// FlashUser is the configuration of [users] section.
// +k8s:openapi-gen=false
type FlashUser struct {
	// +optional
	Readonly *User `json:"readonly,omitempty" toml:"readonly,omitempty"`
	Default  *User `json:"default,omitempty" toml:"default,omitempty"`
}

// User is the configuration of users.
// +k8s:openapi-gen=false
type User struct {
	// +optional
	Password string `json:"password,omitempty" toml:"password"`
	// +optional
	Profile *string `json:"profile,omitempty" toml:"profile,omitempty"`
	// +optional
	Quota *string `json:"quota,omitempty" toml:"quota,omitempty"`
	// +optional
	Networks *Networks `json:"networks,omitempty" toml:"networks,omitempty"`
}

// Networks is the configuration of [users.readonly.networks] section.
// +k8s:openapi-gen=false
type Networks struct {
	// +optional
	IP *string `json:"ip,omitempty" toml:"ip,omitempty"`
}

// FlashQuota is the configuration of [quotas] section.
// +k8s:openapi-gen=false
type FlashQuota struct {
	// +optional
	Default *Quota `json:"default,omitempty" toml:"default,omitempty"`
}

// Quota is the configuration of [quotas.default] section.
// +k8s:openapi-gen=false
type Quota struct {
	// +optional
	Interval *Interval `json:"interval,omitempty" toml:"interval,omitempty"`
}

// Interval is the configuration of [quotas.default.interval] section.
// +k8s:openapi-gen=false
type Interval struct {
	// Optional: Defaults to 3600
	// +optional
	Duration *int32 `json:"duration,omitempty" toml:"duration,omitempty"`
	// Optional: Defaults to 0
	// +optional
	Queries *int32 `json:"queries,omitempty" toml:"queries,omitempty"`