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

package v1alpha1

// Port from TiKV v3.0.6

// TiKVConfig is the configuration of TiKV.
// +k8s:openapi-gen=true
type TiKVConfig struct {
	// Optional: Defaults to info
	// +optional
	LogLevel *string `json:"log-level,omitempty" toml:"log-level,omitempty"`
	// +optional
	LogFile *string `json:"log-file,omitempty" toml:"log-file,omitempty"`
	// +optional
	LogFormat *string `json:"log-format,omitempty" toml:"log-format,omitempty"`
	// +optional
	SlowLogFile *string `json:"slow-log-file,omitempty" toml:"slow-log-file,omitempty"`
	// +optional
	SlowLogThreshold *string `json:"slow-log-threshold,omitempty" toml:"slow-log-threshold,omitempty"`
	// Optional: Defaults to 24h
	// +optional
	LogRotationTimespan *string `json:"log-rotation-timespan,omitempty" toml:"log-rotation-timespan,omitempty"`
	// +optional
	LogRotationSize *string `json:"log-rotation-size,omitempty" toml:"log-rotation-size,omitempty"`
	// +optional
	RefreshConfigInterval *string `json:"refresh-config-interval,omitempty" toml:"refresh-config-interval,omitempty"`
	// +optional
	PanicWhenUnexpectedKeyOrData *bool `json:"panic-when-unexpected-key-or-data,omitempty" toml:"panic-when-unexpected-key-or-data,omitempty"`
	// +optional
	Server *TiKVServerConfig `json:"server,omitempty" toml:"server,omitempty"`
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
	// +optional
	TiKVPessimisticTxn *TiKVPessimisticTxn `json:"pessimistic-txn,omitempty" toml:"pessimistic-txn,omitempty"`
	// +optional
	Backup *TiKVBackupConfig `json:"backup,omitempty" toml:"backup,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVReadPoolConfig struct {
	// +optional
	Unified *TiKVUnifiedReadPoolConfig `json:"unified,omitempty" toml:"unified,omitempty"`
	// +optional
	Coprocessor *TiKVCoprocessorReadPoolConfig `json:"coprocessor,omitempty" toml:"coprocessor,omitempty"`
	// +optional
	Storage *TiKVStorageReadPoolConfig `json:"storage,omitempty" toml:"storage,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVUnifiedReadPoolConfig struct {
	// +optional
	MinThreadCount *int32 `json:"min-thread-count,omitempty" toml:"min-thread-count,omitempty"`
	// +optional
	MaxThreadCount *int32 `json:"max-thread-count,omitempty" toml:"max-thread-count,omitempty"`
	// Deprecated in v4.0.0
	// +optional
	StackSize *string `json:"stack-size,omitempty" toml:"stack-size,omitempty"`
	// +optional
	MaxTasksPerWorker *int32 `json:"max-tasks-per-worker,omitempty" toml:"max-tasks-per-worker,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVStorageReadPoolConfig struct {
	// Optional: Defaults to 4
	// +optional
	HighConcurrency *int64 `json:"high-concurrency,omitempty" toml:"high-concurrency,omitempty"`
	// Optional: Defaults to 4
	// +optional
	NormalConcurrency *int64 `json:"normal-concurrency,omitempty" toml:"normal-concurrency,omitempty"`
	// Optional: Defaults to 4
	// +optional
	LowConcurrency *int64 `json:"low-concurrency,omitempty" toml:"low-concurrency,omitempty"`
	// Optional: Defaults to 2000
	// +optional
	MaxTasksPerWorkerHigh *int64 `json:"max-tasks-per-worker-high,omitempty" toml:"max-tasks-per-worker-high,omitempty"`
	// Optional: Defaults to 2000
	// +optional
	MaxTasksPerWorkerNormal *int64 `json:"max-tasks-per-worker-normal,omitempty" toml:"max-tasks-per-worker-normal,omitempty"`
	// Optional: Defaults to 2000
	// +optional
	MaxTasksPerWorkerLow *int64 `json:"max-tasks-per-worker-low,omitempty" toml:"max-tasks-per-worker-low,omitempty"`
	// Optional: Defaults to 10MB
	// +optional
	StackSize *string `json:"stack-size,omitempty" toml:"stack-size,omitempty"`
	// Optional: Defaults to true
	// +optional
	UseUnifiedPool *bool `json:"use-unified-pool,omitempty" toml:"use-unified-pool,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVCoprocessorReadPoolConfig struct {
	// Optional: Defaults to 8
	// +optional
	HighConcurrency *int64 `json:"high-concurrency,omitempty" toml:"high-concurrency,omitempty"`
	// Optional: Defaults to 8
	// +optional
	NormalConcurrency *int64 `json:"normal-concurrency,omitempty" toml:"normal-concurrency,omitempty"`
	// Optional: Defaults to 8
	// +optional
	LowConcurrency *int64 `json:"low-concurrency,omitempty" toml:"low-concurrency,omitempty"`
	// Optional: Defaults to 2000
	// +optional
	MaxTasksPerWorkerHigh *int64 `json:"max-tasks-per-worker-high,omitempty" toml:"max-tasks-per-worker-high,omitempty"`
	// Optional: Defaults to 2000
	// +optional
	MaxTasksPerWorkerNormal *int64 `json:"max-tasks-per-worker-normal,omitempty" toml:"max-tasks-per-worker-normal,omitempty"`
	// Optional: Defaults to 2000
	// +optional
	MaxTasksPerWorkerLow *int64 `json:"max-tasks-per-worker-low,omitempty" toml:"max-tasks-per-worker-low,omitempty"`
	// Optional: Defaults to 10MB
	// +optional
	StackSize *string `json:"stack-size,omitempty" toml:"stack-size,omitempty"`
	// Optional: Defaults to true
	// +optional
	UseUnifiedPool *bool `json:"use-unified-pool,omitempty" toml:"use-unified-pool,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVPDConfig struct {
	// The PD endpoints for the client.
	//
	// Default is empty.
	// +optional
	Endpoints []string `json:"endpoints,omitempty" toml:"endpoints,omitempty"`
	// The interval at which to retry a PD connection initialization.
	//
	// Default is 300ms.
	// Optional: Defaults to 300ms
	// +optional
	RetryInterval *string `json:"retry-interval,omitempty" toml:"retry-interval,omitempty"`
	// The maximum number of times to retry a PD connection initialization.
	//
	// Default is isize::MAX, represented by -1.
	// Optional: Defaults to -1
	// +optional
	RetryMaxCount *int64 `json:"retry-max-count,omitempty" toml:"retry-max-count,omitempty"`
	// If the client observes the same error message on retry, it can repeat the message only
	// every `n` times.
	//
	// Default is 10. Set to 1 to disable this feature.
	// Optional: Defaults to 10
	// +optional
	RetryLogEvery *int64 `json:"retry-log-every,omitempty" toml:"retry-log-every,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVRaftDBConfig struct {
	// +optional
	WalRecoveryMode *string `json:"wal-recovery-mode,omitempty" toml:"wal-recovery-mode,omitempty"`
	// +optional
	WalDir *string `json:"wal-dir,omitempty" toml:"wal-dir,omitempty"`
	// +optional
	WalTtlSeconds *int64 `json:"wal-ttl-seconds,omitempty" toml:"wal-ttl-seconds,omitempty"`
	// +optional
	WalSizeLimit *string `json:"wal-size-limit,omitempty" toml:"wal-size-limit,omitempty"`
	// +optional
	MaxTotalWalSize *string `json:"max-total-wal-size,omitempty" toml:"max-total-wal-size,omitempty"`
	// +optional
	MaxBackgroundJobs *int64 `json:"max-background-jobs,omitempty" toml:"max-background-jobs,omitempty"`
	// +optional
	MaxManifestFileSize *string `json:"max-manifest-file-size,omitempty" toml:"max-manifest-file-size,omitempty"`
	// +optional
	CreateIfMissing *bool `json:"create-if-missing,omitempty" toml:"create-if-missing,omitempty"`
	// +optional
	MaxOpenFiles *int64 `json:"max-open-files,omitempty" toml:"max-open-files,omitempty"`
	// +optional
	EnableStatistics *bool `json:"enable-statistics,omitempty" toml:"enable-statistics,omitempty"`
	// +optional
	StatsDumpPeriod *string `json:"stats-dump-period,omitempty" toml:"stats-dump-period,omitempty"`
	// +optional
	CompactionReadaheadSize *string `json:"compaction-readahead-size,omitempty" toml:"compaction-readahead-size,omitempty"`
	// +optional
	InfoLogMaxSize *string `json:"info-log-max-size,omitempty" toml:"info-log-max-size,omitempty"`
	// +optional
	FnfoLogRollTime *string `json:"info-log-roll-time,omitempty" toml:"info-log-roll-time,omitempty"`
	// +optional
	InfoLogKeepLogFileNum *int64 `json:"info-log-keep-log-file-num,omitempty" toml:"info-log-keep-log-file-num,omitempty"`
	// +optional
	InfoLogDir *string `json:"info-log-dir,omitempty" toml:"info-log-dir,omitempty"`
	// +optional
	MaxSubCompactions *int64 `json:"max-sub-compactions,omitempty" toml:"max-sub-compactions,omitempty"`
	// +optional
	WritableFileMaxBufferSize *string `json:"writable-file-max-buffer-size,omitempty" toml:"writable-file-max-buffer-size,omitempty"`
	// +optional
	UseDirectIoForFlushAndCompaction *bool `json:"use-direct-io-for-flush-and-compaction,omitempty" toml:"use-direct-io-for-flush-and-compaction,omitempty"`
	// +optional
	EnablePipelinedWrite *bool `json:"enable-pipelined-write,omitempty" toml:"enable-pipelined-write,omitempty"`
	// +optional
	AllowConcurrentMemtableWrite *bool `json:"allow-concurrent-memtable-write,omitempty" toml:"allow-concurrent-memtable-write,omitempty"`
	// +optional
	BytesPerSync *string `json:"bytes-per-sync,omitempty" toml:"bytes-per-sync,omitempty"`
	// +optional
	WalBytesPerSync *string `json:"wal-bytes-per-sync,omitempty" toml:"wal-bytes-per-sync,omitempty"`
	// +optional
	Defaultcf *TiKVCfConfig `json:"defaultcf,omitempty" toml:"defaultcf,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVSecurityConfig struct {
	// +optional
	CAPath *string `json:"ca-path,omitempty" toml:"ca-path,omitempty"`
	// +optional
	CertPath *string `json:"cert-path,omitempty" toml:"cert-path,omitempty"`
	// +optional
	KeyPath *string `json:"key-path,omitempty" toml:"key-path,omitempty"`
	// CertAllowedCN is the Common Name that allowed
	// +optional
	// +k8s:openapi-gen=false
	CertAllowedCN []string `json:"cert-allowed-cn,omitempty" toml:"cert-allowed-cn,omitempty"`
	// +optional
	OverrideSslTarget *string `json:"override-ssl-target,omitempty" toml:"override-ssl-target,omitempty"`
	// +optional
	CipherFile *string `json:"cipher-file,omitempty" toml:"cipher-file,omitempty"`
	// +optional
	Encryption *TiKVSecurityConfigEncryption `json:"encryption,omitempty" toml:"encryption,omitempty"`
}

type TiKVSecurityConfigEncryption struct {
	// Encryption method to use for data files.
	// Possible values are "plaintext", "aes128-ctr", "aes192-ctr" and "aes256-ctr". Value other than
	// "plaintext" means encryption is enabled, in which case master key must be specified.
	// +optional
	DataEncryptionMethod *string `json:"data-encryption-method,omitempty" toml:"data-encryption-method,omitempty"`
	// Specifies how often TiKV rotates data encryption key.
	// +optional
	DataKeyRotationPeriod *string `json:"data-key-rotation-period,omitempty" toml:"data-key-rotation-period,omitempty"`
	// Specifies master key if encryption is enabled. There are three types of master key:
	//
	//   * "plaintext":
	//
	//     Plaintext as master key means no master key is given and only applicable when
	//     encryption is not enabled, i.e. data-encryption-method = "plaintext". This type doesn't
	//     have sub-config items. Example:
	//
	//     [security.encryption.master-key]
	//     type = "plaintext"
	//
	//   * "kms":
	//
	//     Use a KMS service to supply master key. Currently only AWS KMS is supported. This type of
	//     master key is recommended for production use. Example:
	//
	//     [security.encryption.master-key]
	//     type = "kms"
	//     ## KMS CMK key id. Must be a valid KMS CMK where the TiKV process has access to.
	//     ## In production is recommended to grant access of the CMK to TiKV using IAM.
	//     key-id = "1234abcd-12ab-34cd-56ef-1234567890ab"
	//     ## AWS region of the KMS CMK.
	//     region = "us-west-2"
	//     ## (Optional) AWS KMS service endpoint. Only required when non-default KMS endpoint is
	//     ## desired.
	//     endpoint = "https://kms.us-west-2.amazonaws.com"
	//
	//   * "file":
	//
	//     Supply a custom encryption key stored in a file. It is recommended NOT to use in production,
	//     as it breaks the purpose of encryption at rest, unless the file is stored in tempfs.
	//     The file must contain a 256-bits (32 bytes, regardless of key length implied by
	//     data-encryption-method) key encoded as hex string and end with newline ("\n"). Example:
	//
	//     [security.encryption.master-key]
	//     type = "file"
	//     path = "/path/to/master/key/file"
	// +optional
	MasterKey *TiKVSecurityConfigEncryptionMasterKey `json:"master-key,omitempty" toml:"master-key,omitempty"`
	// Specifies the old master key when rotating master key. Same config format as master-key.
	// The key is only access once during TiKV startup, after that TiKV do not need access to the key.
	// And it is okay to leave the stale previous-master-key config after master key rotation.
	// +optional
	PreviousMasterKey *TiKVSecurityConfigEncryptionPreviousMasterKey `json:"previous-master-key,omitempty" toml:"previous-master-key,omitempty"`
}

type TiKVSecurityConfigEncryptionMasterKey struct {
	// +optional
	Type *string `json:"type" toml:"type,omitempty"`

	// Master key file config
	// If the type set to file, this config should be filled
	MasterKeyFileConfig `json:",inline"`

	// Master key KMS config
	// If the type set to kms, this config should be filled
	MasterKeyKMSConfig `json:",inline"`
}

type TiKVSecurityConfigEncryptionPreviousMasterKey struct {
	// +optional
	Type *string `json:"type" toml:"type,omitempty"`

	// Master key file config
	// If the type set to file, this config should be filled
	MasterKeyFileConfig `json:",inline"`

	// Master key KMS config
	// If the type set to kms, this config should be filled
	MasterKeyKMSConfig `json:",inline"`
}

// +k8s:openapi-gen=true
type TiKVImportConfig struct {
	// +optional
	ImportDir *string `json:"import-dir,omitempty" toml:"import-dir,omitempty"`
	// +optional
	NumThreads *int64 `json:"num-threads,omitempty" toml:"num-threads,omitempty"`
	// +optional
	NumImportJobs *int64 `json:"num-import-jobs,omitempty" toml:"num-import-jobs,omitempty"`
	// +optional
	NumImportSstJobs *int64 `json:"num-import-sst-jobs,omitempty" toml:"num-import-sst-jobs,omitempty"`
	// +optional
	MaxPrepareDuration *string `json:"max-prepare-duration,omitempty" toml:"max-prepare-duration,omitempty"`
	// +optional
	RegionSplitSize *string `json:"region-split-size,omitempty" toml:"region-split-size,omitempty"`
	// +optional
	StreamChannelWindow *int64 `json:"stream-channel-window,omitempty" toml:"stream-channel-window,omitempty"`
	// +optional
	MaxOpenEngines *int64 `json:"max-open-engines,omitempty" toml:"max-open-engines,omitempty"`
	// +optional
	UploadSpeedLimit *string `json:"upload-speed-limit,omitempty" toml:"upload-speed-limit,omitempty"`
}

// +k8s:openapi-gen=true
type TiKVGCConfig struct {
	// +optional
	// Optional: Defaults to 512
	BatchKeys *int64 `json:"batch-keys,omitempty" toml:"batch-keys,omitempty"`
	// +optional
	MaxWriteBytesPerSec *string `json:"max-write-bytes-per-sec,omitempty" toml:"max-write-bytes-per-sec,omitempty"`
	// +optional
	EnableCompactionFilter *bool `json:"enable-compaction-filter,omitempty" toml:"enable-compaction-filter,omitempty"`
	// +optional
	EnableCompactionFilterSkipVersionCheck *bool `json:"compaction-filter-skip-version-check,omitempty" toml:"compaction-filter-skip-version-check,omitempty"`
}

// TiKVDbConfig is the rocksdb config.
// +k8s:openapi-gen=true
type TiKVDbConfig struct {
	// +optional
	// Optional: Defaults to 2
	WalRecoveryMode *int64 `json:"wal-recovery-mode,omitempty" toml:"wal-recovery-mode,omitempty"`
	// +optional
	WalTTLSeconds *int64 `json:"wal-ttl-seconds,omitempty" toml:"wal-ttl-seconds,omitempty"`
	// +optional
	WalSizeLimit *string `json:"wal-size-limit,omitempty" toml:"wal-size-limit,omitempty"`
	// +optional
	// Optional: Defaults to 4GB
	MaxTotalWalSize *string `json:"max-total-wal-size,omitempty" toml:"max-total-wal-size,omitempty"`
	// +optional
	// Optional: Defaults to 8
	MaxBackgroundJobs *int64 `json:"max-background-jobs,omitempty" toml:"max-background-jobs,omitempty"`
	// +optional
	// Optional: Defaults to 128MB
	MaxManifestFileSize *string `json:"max-manifest-file-size,omitempty" toml:"max-manifest-file-size,omitempty"`
	// +optional
	// Optional: Defaults to true
	CreateIfMissing *bool `json:"create-if-missing,omitempty" toml:"create-if-missing,omitempty"`
	// +optional
	// Optional: Defaults to 40960
	MaxOpenFiles *int64 `json:"max-open-files,omitempty" toml:"max-open-files,omitempty"`
	// +optional
	// Optional: Defaults to true
	EnableStatistics *bool `json:"enable-statistics,omitempty" toml:"enable-statistics,omitempty"`
	// +optional
	// Optional: Defaults to 10m
	StatsDumpPeriod *string `json:"stats-dump-period,omitempty" toml:"stats-dump-period,omitempty"`
	// Optional: Defaults to 0
	// +optional
	CompactionReadaheadSize *string `json:"compaction-readahead-size,omitempty" toml:"compaction-readahead-size,omitempty"`
	// +optional
	InfoLogMaxSize *string `json:"info-log-max-size,omitempty" toml:"info-log-max-size,omitempty"`
	// +optional
	InfoLogRollTime *string `json:"info-log-roll-time,omitempty" toml:"info-log-roll-time,omitempty"`
	// +optional
	InfoLogKeepLogFileNum *int64 `json:"info-log-keep-log-file-num,omitempty" toml:"info-log-keep-log-file-num,omitempty"`
	// +optional
	InfoLogDir 