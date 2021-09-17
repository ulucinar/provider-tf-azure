// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtection) DeepCopyInto(out *AdvancedThreatProtection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtection.
func (in *AdvancedThreatProtection) DeepCopy() *AdvancedThreatProtection {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdvancedThreatProtection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionList) DeepCopyInto(out *AdvancedThreatProtectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdvancedThreatProtection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionList.
func (in *AdvancedThreatProtectionList) DeepCopy() *AdvancedThreatProtectionList {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdvancedThreatProtectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionObservation) DeepCopyInto(out *AdvancedThreatProtectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionObservation.
func (in *AdvancedThreatProtectionObservation) DeepCopy() *AdvancedThreatProtectionObservation {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionParameters) DeepCopyInto(out *AdvancedThreatProtectionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionParameters.
func (in *AdvancedThreatProtectionParameters) DeepCopy() *AdvancedThreatProtectionParameters {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionSpec) DeepCopyInto(out *AdvancedThreatProtectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionSpec.
func (in *AdvancedThreatProtectionSpec) DeepCopy() *AdvancedThreatProtectionSpec {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionStatus) DeepCopyInto(out *AdvancedThreatProtectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionStatus.
func (in *AdvancedThreatProtectionStatus) DeepCopy() *AdvancedThreatProtectionStatus {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionStatus)
	in.DeepCopyInto(out)
	return out
}
