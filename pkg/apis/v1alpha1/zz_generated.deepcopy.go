// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/kedacore/keda/v2/api/v1alpha1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildImpl) DeepCopyInto(out *BuildImpl) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SrcRepo != nil {
		in, out := &in.SrcRepo, &out.SrcRepo
		*out = new(GitRepo)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(Registry)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildImpl.
func (in *BuildImpl) DeepCopy() *BuildImpl {
	if in == nil {
		return nil
	}
	out := new(BuildImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Builder) DeepCopyInto(out *Builder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Builder.
func (in *Builder) DeepCopy() *Builder {
	if in == nil {
		return nil
	}
	out := new(Builder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Builder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderList) DeepCopyInto(out *BuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Builder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderList.
func (in *BuilderList) DeepCopy() *BuilderList {
	if in == nil {
		return nil
	}
	out := new(BuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderSpec) DeepCopyInto(out *BuilderSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SrcRepo != nil {
		in, out := &in.SrcRepo, &out.SrcRepo
		*out = new(GitRepo)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(Registry)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderSpec.
func (in *BuilderSpec) DeepCopy() *BuilderSpec {
	if in == nil {
		return nil
	}
	out := new(BuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderStatus) DeepCopyInto(out *BuilderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderStatus.
func (in *BuilderStatus) DeepCopy() *BuilderStatus {
	if in == nil {
		return nil
	}
	out := new(BuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaprComponent) DeepCopyInto(out *DaprComponent) {
	*out = *in
	in.ComponentSpec.DeepCopyInto(&out.ComponentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaprComponent.
func (in *DaprComponent) DeepCopy() *DaprComponent {
	if in == nil {
		return nil
	}
	out := new(DaprComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaprInput) DeepCopyInto(out *DaprInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaprInput.
func (in *DaprInput) DeepCopy() *DaprInput {
	if in == nil {
		return nil
	}
	out := new(DaprInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaprOutput) DeepCopyInto(out *DaprOutput) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaprOutput.
func (in *DaprOutput) DeepCopy() *DaprOutput {
	if in == nil {
		return nil
	}
	out := new(DaprOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaprRuntime) DeepCopyInto(out *DaprRuntime) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]DaprComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subscriptions != nil {
		in, out := &in.Subscriptions, &out.Subscriptions
		*out = make([]DaprSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = new(DaprInput)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = make([]DaprOutput, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScaledObject != nil {
		in, out := &in.ScaledObject, &out.ScaledObject
		*out = new(KedaScaledObject)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaledJob != nil {
		in, out := &in.ScaledJob, &out.ScaledJob
		*out = new(KedaScaledJob)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaprRuntime.
func (in *DaprRuntime) DeepCopy() *DaprRuntime {
	if in == nil {
		return nil
	}
	out := new(DaprRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaprSubscription) DeepCopyInto(out *DaprSubscription) {
	*out = *in
	out.SubscriptionSpec = in.SubscriptionSpec
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaprSubscription.
func (in *DaprSubscription) DeepCopy() *DaprSubscription {
	if in == nil {
		return nil
	}
	out := new(DaprSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(BuildImpl)
		(*in).DeepCopyInto(*out)
	}
	if in.Serving != nil {
		in, out := &in.Serving, &out.Serving
		*out = new(ServingImpl)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepo) DeepCopyInto(out *GitRepo) {
	*out = *in
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(string)
		**out = **in
	}
	if in.Refspec != nil {
		in, out := &in.Refspec, &out.Refspec
		*out = new(string)
		**out = **in
	}
	if in.Submodules != nil {
		in, out := &in.Submodules, &out.Submodules
		*out = new(bool)
		**out = **in
	}
	if in.Depth != nil {
		in, out := &in.Depth, &out.Depth
		*out = new(int8)
		**out = **in
	}
	if in.SslVerify != nil {
		in, out := &in.SslVerify, &out.SslVerify
		*out = new(bool)
		**out = **in
	}
	if in.SubDirectory != nil {
		in, out := &in.SubDirectory, &out.SubDirectory
		*out = new(string)
		**out = **in
	}
	if in.SourceSubPath != nil {
		in, out := &in.SourceSubPath, &out.SourceSubPath
		*out = new(string)
		**out = **in
	}
	if in.DeleteExisting != nil {
		in, out := &in.DeleteExisting, &out.DeleteExisting
		*out = new(string)
		**out = **in
	}
	if in.HttpProxy != nil {
		in, out := &in.HttpProxy, &out.HttpProxy
		*out = new(string)
		**out = **in
	}
	if in.HttpsProxy != nil {
		in, out := &in.HttpsProxy, &out.HttpsProxy
		*out = new(string)
		**out = **in
	}
	if in.NoProxy != nil {
		in, out := &in.NoProxy, &out.NoProxy
		*out = new(string)
		**out = **in
	}
	if in.Verbose != nil {
		in, out := &in.Verbose, &out.Verbose
		*out = new(bool)
		**out = **in
	}
	if in.GitInitImage != nil {
		in, out := &in.GitInitImage, &out.GitInitImage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepo.
func (in *GitRepo) DeepCopy() *GitRepo {
	if in == nil {
		return nil
	}
	out := new(GitRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaScaledJob) DeepCopyInto(out *KedaScaledJob) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	in.ScalingStrategy.DeepCopyInto(&out.ScalingStrategy)
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]apiv1alpha1.ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaScaledJob.
func (in *KedaScaledJob) DeepCopy() *KedaScaledJob {
	if in == nil {
		return nil
	}
	out := new(KedaScaledJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaScaledObject) DeepCopyInto(out *KedaScaledObject) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(apiv1alpha1.AdvancedConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]apiv1alpha1.ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaScaledObject.
func (in *KedaScaledObject) DeepCopy() *KedaScaledObject {
	if in == nil {
		return nil
	}
	out := new(KedaScaledObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.Account != nil {
		in, out := &in.Account, &out.Account
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Serving) DeepCopyInto(out *Serving) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Serving.
func (in *Serving) DeepCopy() *Serving {
	if in == nil {
		return nil
	}
	out := new(Serving)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Serving) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingImpl) DeepCopyInto(out *ServingImpl) {
	*out = *in
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(Runtime)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Dapr != nil {
		in, out := &in.Dapr, &out.Dapr
		*out = new(DaprRuntime)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingImpl.
func (in *ServingImpl) DeepCopy() *ServingImpl {
	if in == nil {
		return nil
	}
	out := new(ServingImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingList) DeepCopyInto(out *ServingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Serving, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingList.
func (in *ServingList) DeepCopy() *ServingList {
	if in == nil {
		return nil
	}
	out := new(ServingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingSpec) DeepCopyInto(out *ServingSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(Runtime)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Dapr != nil {
		in, out := &in.Dapr, &out.Dapr
		*out = new(DaprRuntime)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingSpec.
func (in *ServingSpec) DeepCopy() *ServingSpec {
	if in == nil {
		return nil
	}
	out := new(ServingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingStatus) DeepCopyInto(out *ServingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingStatus.
func (in *ServingStatus) DeepCopy() *ServingStatus {
	if in == nil {
		return nil
	}
	out := new(ServingStatus)
	in.DeepCopyInto(out)
	return out
}
