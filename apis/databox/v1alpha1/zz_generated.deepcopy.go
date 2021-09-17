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
func (in *ContactObservation) DeepCopyInto(out *ContactObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactObservation.
func (in *ContactObservation) DeepCopy() *ContactObservation {
	if in == nil {
		return nil
	}
	out := new(ContactObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactParameters) DeepCopyInto(out *ContactParameters) {
	*out = *in
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactParameters.
func (in *ContactParameters) DeepCopy() *ContactParameters {
	if in == nil {
		return nil
	}
	out := new(ContactParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeDevice) DeepCopyInto(out *DataboxEdgeDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeDevice.
func (in *DataboxEdgeDevice) DeepCopy() *DataboxEdgeDevice {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataboxEdgeDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeDeviceList) DeepCopyInto(out *DataboxEdgeDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataboxEdgeDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeDeviceList.
func (in *DataboxEdgeDeviceList) DeepCopy() *DataboxEdgeDeviceList {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataboxEdgeDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeDeviceObservation) DeepCopyInto(out *DataboxEdgeDeviceObservation) {
	*out = *in
	if in.DeviceProperties != nil {
		in, out := &in.DeviceProperties, &out.DeviceProperties
		*out = make([]DevicePropertiesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeDeviceObservation.
func (in *DataboxEdgeDeviceObservation) DeepCopy() *DataboxEdgeDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeDeviceParameters) DeepCopyInto(out *DataboxEdgeDeviceParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeDeviceParameters.
func (in *DataboxEdgeDeviceParameters) DeepCopy() *DataboxEdgeDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeDeviceSpec) DeepCopyInto(out *DataboxEdgeDeviceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeDeviceSpec.
func (in *DataboxEdgeDeviceSpec) DeepCopy() *DataboxEdgeDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeDeviceStatus) DeepCopyInto(out *DataboxEdgeDeviceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeDeviceStatus.
func (in *DataboxEdgeDeviceStatus) DeepCopy() *DataboxEdgeDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeOrder) DeepCopyInto(out *DataboxEdgeOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeOrder.
func (in *DataboxEdgeOrder) DeepCopy() *DataboxEdgeOrder {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataboxEdgeOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeOrderList) DeepCopyInto(out *DataboxEdgeOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataboxEdgeOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeOrderList.
func (in *DataboxEdgeOrderList) DeepCopy() *DataboxEdgeOrderList {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataboxEdgeOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeOrderObservation) DeepCopyInto(out *DataboxEdgeOrderObservation) {
	*out = *in
	if in.ReturnTracking != nil {
		in, out := &in.ReturnTracking, &out.ReturnTracking
		*out = make([]ReturnTrackingObservation, len(*in))
		copy(*out, *in)
	}
	if in.ShipmentHistory != nil {
		in, out := &in.ShipmentHistory, &out.ShipmentHistory
		*out = make([]ShipmentHistoryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShipmentTracking != nil {
		in, out := &in.ShipmentTracking, &out.ShipmentTracking
		*out = make([]ShipmentTrackingObservation, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]StatusObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeOrderObservation.
func (in *DataboxEdgeOrderObservation) DeepCopy() *DataboxEdgeOrderObservation {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeOrderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeOrderParameters) DeepCopyInto(out *DataboxEdgeOrderParameters) {
	*out = *in
	if in.Contact != nil {
		in, out := &in.Contact, &out.Contact
		*out = make([]ContactParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShipmentAddress != nil {
		in, out := &in.ShipmentAddress, &out.ShipmentAddress
		*out = make([]ShipmentAddressParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeOrderParameters.
func (in *DataboxEdgeOrderParameters) DeepCopy() *DataboxEdgeOrderParameters {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeOrderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeOrderSpec) DeepCopyInto(out *DataboxEdgeOrderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeOrderSpec.
func (in *DataboxEdgeOrderSpec) DeepCopy() *DataboxEdgeOrderSpec {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataboxEdgeOrderStatus) DeepCopyInto(out *DataboxEdgeOrderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataboxEdgeOrderStatus.
func (in *DataboxEdgeOrderStatus) DeepCopy() *DataboxEdgeOrderStatus {
	if in == nil {
		return nil
	}
	out := new(DataboxEdgeOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePropertiesObservation) DeepCopyInto(out *DevicePropertiesObservation) {
	*out = *in
	if in.ConfiguredRoleTypes != nil {
		in, out := &in.ConfiguredRoleTypes, &out.ConfiguredRoleTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePropertiesObservation.
func (in *DevicePropertiesObservation) DeepCopy() *DevicePropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(DevicePropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePropertiesParameters) DeepCopyInto(out *DevicePropertiesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePropertiesParameters.
func (in *DevicePropertiesParameters) DeepCopy() *DevicePropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(DevicePropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReturnTrackingObservation) DeepCopyInto(out *ReturnTrackingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReturnTrackingObservation.
func (in *ReturnTrackingObservation) DeepCopy() *ReturnTrackingObservation {
	if in == nil {
		return nil
	}
	out := new(ReturnTrackingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReturnTrackingParameters) DeepCopyInto(out *ReturnTrackingParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReturnTrackingParameters.
func (in *ReturnTrackingParameters) DeepCopy() *ReturnTrackingParameters {
	if in == nil {
		return nil
	}
	out := new(ReturnTrackingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentAddressObservation) DeepCopyInto(out *ShipmentAddressObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentAddressObservation.
func (in *ShipmentAddressObservation) DeepCopy() *ShipmentAddressObservation {
	if in == nil {
		return nil
	}
	out := new(ShipmentAddressObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentAddressParameters) DeepCopyInto(out *ShipmentAddressParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentAddressParameters.
func (in *ShipmentAddressParameters) DeepCopy() *ShipmentAddressParameters {
	if in == nil {
		return nil
	}
	out := new(ShipmentAddressParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentHistoryObservation) DeepCopyInto(out *ShipmentHistoryObservation) {
	*out = *in
	if in.AdditionalDetails != nil {
		in, out := &in.AdditionalDetails, &out.AdditionalDetails
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentHistoryObservation.
func (in *ShipmentHistoryObservation) DeepCopy() *ShipmentHistoryObservation {
	if in == nil {
		return nil
	}
	out := new(ShipmentHistoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentHistoryParameters) DeepCopyInto(out *ShipmentHistoryParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentHistoryParameters.
func (in *ShipmentHistoryParameters) DeepCopy() *ShipmentHistoryParameters {
	if in == nil {
		return nil
	}
	out := new(ShipmentHistoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentTrackingObservation) DeepCopyInto(out *ShipmentTrackingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentTrackingObservation.
func (in *ShipmentTrackingObservation) DeepCopy() *ShipmentTrackingObservation {
	if in == nil {
		return nil
	}
	out := new(ShipmentTrackingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipmentTrackingParameters) DeepCopyInto(out *ShipmentTrackingParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipmentTrackingParameters.
func (in *ShipmentTrackingParameters) DeepCopy() *ShipmentTrackingParameters {
	if in == nil {
		return nil
	}
	out := new(ShipmentTrackingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusObservation) DeepCopyInto(out *StatusObservation) {
	*out = *in
	if in.AdditionalDetails != nil {
		in, out := &in.AdditionalDetails, &out.AdditionalDetails
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusObservation.
func (in *StatusObservation) DeepCopy() *StatusObservation {
	if in == nil {
		return nil
	}
	out := new(StatusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusParameters) DeepCopyInto(out *StatusParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusParameters.
func (in *StatusParameters) DeepCopy() *StatusParameters {
	if in == nil {
		return nil
	}
	out := new(StatusParameters)
	in.DeepCopyInto(out)
	return out
}
