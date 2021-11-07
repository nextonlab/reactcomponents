// +build !ignore_autogenerated

// Copyright PingCAP, Inc.
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	model "github.com/prometheus/common/model"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"

	binlog "github.com/pingcap/tipocket/pkg/tidb-operator/binlog"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoResource) DeepCopyInto(out *AutoResource) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	out.Storage = in.Storage.DeepCopy()
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoResource.
func (in *AutoResource) DeepCopy() *AutoResource {
	if in == nil {
		return nil
	}
	out := new(AutoResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoRule) DeepCopyInto(out *AutoRule) {
	*out = *in
	if in.MinThreshold != nil {
		in, out := &in.MinThreshold, &out.MinThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ResourceTypes != nil {
		in, out := &in.ResourceTypes, &out.ResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoRule.
func (in *AutoRule) DeepCopy() *AutoRule {
	if in == nil {
		return nil
	}
	out := new(AutoRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BRConfig) DeepCopyInto(out *BRConfig) {
	*out = *in
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(uint32)
		**out = **in
	}
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(uint)
		**out = **in
	}
	if in.Checksum != nil {
		in, out := &in.Checksum, &out.Checksum
		*out = new(bool)
		**out = **in
	}
	if in.SendCredToTikv != nil {
		in, out := &in.SendCredToTikv, &out.SendCredToTikv
		*out = new(bool)
		**out = **in
	}
	if in.OnLine != nil {
		in, out := &in.OnLine, &out.OnLine
		*out = new(bool)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BRConfig.
func (in *BRConfig) DeepCopy() *BRConfig {
	if in == nil {
		return nil
	}
	out := new(BRConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupCondition) DeepCopyInto(out *BackupCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupCondition.
func (in *BackupCondition) DeepCopy() *BackupCondition {
	if in == nil {
		return nil
	}
	out := new(BackupCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSchedule) DeepCopyInto(out *BackupSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSchedule.
func (in *BackupSchedule) DeepCopy() *BackupSchedule {
	if in == nil {
		return nil
	}
	out := new(BackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleList) DeepCopyInto(out *BackupScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleList.
func (in *BackupScheduleList) DeepCopy() *BackupScheduleList {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleSpec) DeepCopyInto(out *BackupScheduleSpec) {
	*out = *in
	if in.MaxBackups != nil {
		in, out := &in.MaxBackups, &out.MaxBackups
		*out = new(int32)
		**out = **in
	}
	if in.MaxReservedTime != nil {
		in, out := &in.MaxReservedTime, &out.MaxReservedTime
		*out = new(string)
		**out = **in
	}
	in.BackupTemplate.DeepCopyInto(&out.BackupTemplate)
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleSpec.
func (in *BackupScheduleSpec) DeepCopy() *BackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleStatus) DeepCopyInto(out *BackupScheduleStatus) {
	*out = *in
	if in.LastBackupTime != nil {
		in, out := &in.LastBackupTime, &out.LastBackupTime
		*out = (*in).DeepCopy()
	}
	if in.AllBackupCleanTime != nil {
		in, out := &in.AllBackupCleanTime, &out.AllBackupCleanTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleStatus.
func (in *BackupScheduleStatus) DeepCopy() *BackupScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(TiDBAccessConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TikvGCLifeTime != nil {
		in, out := &in.TikvGCLifeTime, &out.TikvGCLifeTime
		*out = new(string)
		**out = **in
	}
	in.StorageProvider.DeepCopyInto(&out.StorageProvider)
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.BR != nil {
		in, out := &in.BR, &out.BR
		*out = new(BRConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Dumpling != nil {
		in, out := &in.Dumpling, &out.Dumpling
		*out = new(DumplingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.TableFilter != nil {
		in, out := &in.TableFilter, &out.TableFilter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	in.TimeStarted.DeepCopyInto(&out.TimeStarted)
	in.TimeCompleted.DeepCopyInto(&out.TimeCompleted)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]BackupCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAutoScalerSpec) DeepCopyInto(out *BasicAutoScalerSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make(map[v1.ResourceName]AutoRule, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ScaleInIntervalSeconds != nil {
		in, out := &in.ScaleInIntervalSeconds, &out.ScaleInIntervalSeconds
		*out = new(int32)
		**out = **in
	}
	if in.ScaleOutIntervalSeconds != nil {
		in, out := &in.ScaleOutIntervalSeconds, &out.ScaleOutIntervalSeconds
		*out = new(int32)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]AutoResource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAutoScalerSpec.
func (in *BasicAutoScalerSpec) DeepCopy() *BasicAutoScalerSpec {
	if in == nil {
		return nil
	}
	out := new(BasicAutoScalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAutoScalerStatus) DeepCopyInto(out *BasicAutoScalerStatus) {
	*out = *in
	if in.LastAutoScalingTimestamp != nil {
		in, out := &in.LastAutoScalingTimestamp, &out.LastAutoScalingTimestamp
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAutoScalerStatus.
func (in *BasicAutoScalerStatus) DeepCopy() *BasicAutoScalerStatus {
	if in == nil {
		return nil
	}
	out := new(BasicAutoScalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binlog) DeepCopyInto(out *Binlog) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.WriteTimeout != nil {
		in, out := &in.WriteTimeout, &out.WriteTimeout
		*out = new(string)
		**out = **in
	}
	if in.IgnoreError != nil {
		in, out := &in.IgnoreError, &out.IgnoreError
		*out = new(bool)
		**out = **in
	}
	if in.BinlogSocket != nil {
		in, out := &in.BinlogSocket, &out.BinlogSocket
		*out = new(string)
		**out = **in
	}
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binlog.
func (in *Binlog) DeepCopy() *Binlog {
	if in == nil {
		return nil
	}
	out := new(Binlog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDCConfigWraper) DeepCopyInto(out *CDCConfigWraper) {
	*out = *in
	if in.GenericConfig != nil {
		in, out := &in.GenericConfig, &out.GenericConfig
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDCConfigWraper.
func (in *CDCConfigWraper) DeepCopy() *CDCConfigWraper {
	if in == nil {
		return nil
	}
	out := new(CDCConfigWraper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRef) DeepCopyInto(out *ClusterRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRef.
func (in *ClusterRef) DeepCopy() *ClusterRef {
	if in == nil {
		return nil
	}
	out := new(ClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonConfig) DeepCopyInto(out *CommonConfig) {
	*out = *in
	if in.TmpPath != nil {
		in, out := &in.TmpPath, &out.TmpPath
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DefaultProfile != nil {
		in, out := &in.DefaultProfile, &out.DefaultProfile
		*out = new(string)
		**out = **in
	}
	if in.FlashDataPath != nil {
		in, out := &in.FlashDataPath, &out.FlashDataPath
		*out = new(string)
		**out = **in
	}
	if in.PathRealtimeMode != nil {
		in, out := &in.PathRealtimeMode, &out.PathRealtimeMode
		*out = new(bool)
		**out = **in
	}
	if in.MarkCacheSize != nil {
		in, out := &in.MarkCacheSize, &out.MarkCacheSize
		*out = new(int64)
		**out = **in
	}
	if in.MinmaxIndexCacheSize != nil {
		in, out := &in.MinmaxIndexCacheSize, &out.MinmaxIndexCacheSize
		*out = new(int64)
		**out = **in
	}
	if in.ListenHost != nil {
		in, out := &in.ListenHost, &out.ListenHost
		*out = new(string)
		**out = **in
	}
	if in.TCPPort != nil {
		in, out := &in.TCPPort, &out.TCPPort
		*out = new(int32)
		**out = **in
	}
	if in.HTTPPort != nil {
		in, out := &in.HTTPPort, &out.HTTPPort
		*out = new(int32)
		**out = **in
	}
	if in.TCPPortSecure != nil {
		in, out := &in.TCPPortSecure, &out.TCPPortSecure
		*out = new(int32)
		**out = **in
	}
	if in.HTTPSPort != nil {
		in, out := &in.HTTPSPort, &out.HTTPSPort
		*out = new(int32)
		**out = **in
	}
	if in.InternalServerHTTPPort != nil {
		in, out := &in.InternalServerHTTPPort, &out.InternalServerHTTPPort
		*out = new(int32)
		**out = **in
	}
	if in.Flash != nil {
		in, out := &in.Flash, &out.Flash
		*out = new(Flash)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashLogger != nil {
		in, out := &in.FlashLogger, &out.FlashLogger
		*out = new(FlashLogger)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashApplication != nil {
		in, out := &in.FlashApplication, &out.FlashApplication
		*out = new(FlashApplication)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashRaft != nil {
		in, out := &in.FlashRaft, &out.FlashRaft
		*out = new(FlashRaft)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashStatus != nil {
		in, out := &in.FlashStatus, &out.FlashStatus
		*out = new(FlashStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashQuota != nil {
		in, out := &in.FlashQuota, &out.FlashQuota
		*out = new(FlashQuota)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashUser != nil {
		in, out := &in.FlashUser, &out.FlashUser
		*out = new(FlashUser)
		(*in).DeepCopyInto(*out)
	}
	if in.FlashProfile != nil {
		in, out := &in.FlashProfile, &out.FlashProfile
		*out = new(FlashProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(FlashSecurity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonConfig.
func (in *CommonConfig) DeepCopy() *CommonConfig {
	if in == nil {
		return nil
	}
	out := new(CommonConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigUpdateStrategy != nil {
		in, out := &in.ConfigUpdateStrategy, &out.ConfigUpdateStrategy
		*out = new(ConfigUpdateStrategy)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalContainers != nil {
		in, out := &in.AdditionalContainers, &out.AdditionalContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalVolumes != nil {
		in, out := &in.AdditionalVolumes, &out.AdditionalVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalVolumeMounts != nil {
		in, out := &in.AdditionalVolumeMounts, &out.AdditionalVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]TopologySpreadConstraint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapRef) DeepCopyInto(out *ConfigMapRef) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapRef.
func (in *ConfigMapRef) DeepCopy() *ConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(ConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoprocessorCache) DeepCopyInto(out *CoprocessorCache) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.CapacityMB != nil {
		in, out := &in.CapacityMB, &out.CapacityMB
		*out = new(float64)
		**out = **in
	}
	if in.AdmissionMaxResultMB != nil {
		in, out := &in.AdmissionMaxResultMB, &out.AdmissionMaxResultMB
		*out = new(float64)
		**out = **in
	}
	if in.AdmissionMinProcessMs != nil {
		in, out := &in.AdmissionMinProcessMs, &out.AdmissionMinProcessMs
		*out = new(uint64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoprocessorCache.
func (in *CoprocessorCache) DeepCopy() *CoprocessorCache {
	if in == nil {
		return nil
	}
	out := new(CoprocessorCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdKind) DeepCopyInto(out *CrdKind) {
	*out = *in
	if in.ShortNames != nil {
		in, out := &in.ShortNames, &out.ShortNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalPrinterColums != nil {
		in, out := &in.AdditionalPrinterColums, &out.AdditionalPrinterColums
		*out = make([]v1beta1.CustomResourceColumnDefinition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdKind.
func (in *CrdKind) DeepCopy() *CrdKind {
	if in == nil {
		return nil
	}
	out := new(CrdKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdKinds) DeepCopyInto(out *CrdKinds) {
	*out = *in
	in.TiDBCluster.DeepCopyInto(&out.TiDBCluster)
	in.DMCluster.DeepCopyInto(&out.DMCluster)
	in.Backup.DeepCopyInto(&out.Backup)
	in.Restore.DeepCopyInto(&out.Restore)
	in.BackupSchedule.DeepCopyInto(&out.BackupSchedule)
	in.TiDBMonitor.DeepCopyInto(&out.TiDBMonitor)
	in.TiDBInitializer.DeepCopyInto(&out.TiDBInitializer)
	in.TidbClusterAutoScaler.DeepCopyInto(&out.TidbClusterAutoScaler)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdKinds.
func (in *CrdKinds) DeepCopy() *CrdKinds {
	if in == nil {
		return nil
	}
	out := new(CrdKinds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMCluster) DeepCopyInto(out *DMCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMCluster.
func (in *DMCluster) DeepCopy() *DMCluster {
	if in == nil {
		return nil
	}
	out := new(DMCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DMCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMClusterCondition) DeepCopyInto(out *DMClusterCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMClusterCondition.
func (in *DMClusterCondition) DeepCopy() *DMClusterCondition {
	if in == nil {
		return nil
	}
	out := new(DMClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMClusterList) DeepCopyInto(out *DMClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DMCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMClusterList.
func (in *DMClusterList) DeepCopy() *DMClusterList {
	if in == nil {
		return nil
	}
	out := new(DMClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DMClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMClusterSpec) DeepCopyInto(out *DMClusterSpec) {
	*out = *in
	in.Discovery.DeepCopyInto(&out.Discovery)
	in.Master.DeepCopyInto(&out.Master)
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = new(WorkerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PVReclaimPolicy != nil {
		in, out := &in.PVReclaimPolicy, &out.PVReclaimPolicy
		*out = new(v1.PersistentVolumeReclaimPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.EnablePVReclaim != nil {
		in, out := &in.EnablePVReclaim, &out.EnablePVReclaim
		*out = new(bool)
		**out = **in
	}
	if in.TLSCluster != nil {
		in, out := &in.TLSCluster, &out.TLSCluster
		*out = new(TLSCluster)
		**out = **in
	}
	if in.TLSClientSecretNames != nil {
		in, out := &in.TLSClientSecretNames, &out.TLSClientSecretNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]TopologySpreadConstraint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMClusterSpec.
func (in *DMClusterSpec) DeepCopy() *DMClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DMClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMClusterStatus) DeepCopyInto(out *DMClusterStatus) {
	*out = *in
	in.Master.DeepCopyInto(&out.Master)
	in.Worker.DeepCopyInto(&out.Worker)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DMClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMClusterStatus.
func (in *DMClusterStatus) DeepCopy() *DMClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DMClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMDiscoverySpec) DeepCopyInto(out *DMDiscoverySpec) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMDiscoverySpec.
func (in *DMDiscoverySpec) DeepCopy() *DMDiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(DMDiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMMonitorSpec) DeepCopyInto(out *DMMonitorSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ClusterRef, len(*in))
		copy(*out, *in)
	}
	in.Initializer.DeepCopyInto(&out.Initializer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMMonitorSpec.
func (in *DMMonitorSpec) DeepCopy() *DMMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(DMMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DMSecurityConfig) DeepCopyInto(out *DMSecurityConfig) {
	*out = *in
	if in.SSLCA != nil {
		in, out := &in.SSLCA, &out.SSLCA
		*out = new(string)
		**out = **in
	}
	if in.SSLCert != nil {
		in, out := &in.SSLCert, &out.SSLCert
		*out = new(string)
		**out = **in
	}
	if in.SSLKey != nil {
		in, out := &in.SSLKey, &out.SSLKey
		*out = new(string)
		**out = **in
	}
	if in.CertAllowedCN != nil {
		in, out := &in.CertAllowedCN, &out.CertAllowedCN
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DMSecurityConfig.
func (in *DMSecurityConfig) DeepCopy() *DMSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(DMSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardConfig) DeepCopyInto(out *DashboardConfig) {
	*out = *in
	if in.TiDBCAPath != nil {
		in, out := &in.TiDBCAPath, &out.TiDBCAPath
		*out = new(string)
		**out = **in
	}
	if in.TiDBCertPath != nil {
		in, out := &in.TiDBCertPath, &out.TiDBCertPath
		*out = new(string)
		**out = **in
	}
	if in.TiDBKeyPath != nil {
		in, out := &in.TiDBKeyPath, &out.TiDBKeyPath
		*out = new(string)
		**out = **in
	}
	if in.PublicPathPrefix != nil {
		in, out := &in.PublicPathPrefix, &out.PublicPathPrefix
		*out = new(string)
		**out = **in
	}
	if in.InternalProxy != nil {
		in, out := &in.InternalProxy, &out.InternalProxy
		*out = new(bool)
		**out = **in
	}
	if in.DisableTelemetry != nil {
		in, out := &in.DisableTelemetry, &out.DisableTelemetry
		*out = new(bool)
		**out = **in
	}
	if in.EnableTelemetry != nil {
		in, out := &in.EnableTelemetry, &out.EnableTelemetry
		*out = new(bool)
		**out = **in
	}
	if in.EnableExperimental != nil {
		in, out := &in.EnableExperimental, &out.EnableExperimental
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardConfig.
func (in *DashboardConfig) DeepCopy() *DashboardConfig {
	if in == nil {
		return nil
	}
	out := new(DashboardConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataResource) DeepCopyInto(out *DataResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataResource.
func (in *DataResource) DeepCopy() *DataResource {
	if in == nil {
		return nil
	}
	out := new(DataResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataResourceList) DeepCopyInto(out *DataResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataResourceList.
func (in *DataResourceList) DeepCopy() *DataResourceList {
	if in == nil {
		return nil
	}
	out := new(DataResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStorageStatus) DeepCopyInto(out *DeploymentStorageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStorageStatus.
func (in *DeploymentStorageStatus) DeepCopy() *DeploymentStorageStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoverySpec) DeepCopyInto(out *DiscoverySpec) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoverySpec.
func (in *DiscoverySpec) DeepCopy() *DiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(DiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DumplingConfig) DeepCopyInto(out *DumplingConfig) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TableFilter != nil {
		in, out := &in.TableFilter, &out.TableFilter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DumplingConfig.
func (in *DumplingConfig) DeepCopy() *DumplingConfig {
	if in == nil {
		return nil
	}
	out := new(DumplingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the 