/*
Copyright 2022 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AllowRuleObservation struct {
}

type AllowRuleParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionFromIpsNotAllowed []*string `json:"connectionFromIpsNotAllowed,omitempty" tf:"connection_from_ips_not_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionToIPNotAllowed []*string `json:"connectionToIpNotAllowed,omitempty" tf:"connection_to_ip_not_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionToIpsNotAllowed []*string `json:"connectionToIpsNotAllowed,omitempty" tf:"connection_to_ips_not_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	LocalUserNotAllowed []*string `json:"localUserNotAllowed,omitempty" tf:"local_user_not_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	LocalUsersNotAllowed []*string `json:"localUsersNotAllowed,omitempty" tf:"local_users_not_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessNotAllowed []*string `json:"processNotAllowed,omitempty" tf:"process_not_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessesNotAllowed []*string `json:"processesNotAllowed,omitempty" tf:"processes_not_allowed,omitempty"`
}

type IOTSecurityDeviceGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTSecurityDeviceGroupParameters struct {

	// +kubebuilder:validation:Optional
	AllowRule []AllowRuleParameters `json:"allowRule,omitempty" tf:"allow_rule,omitempty"`

	// +kubebuilder:validation:Required
	IOTHubID *string `json:"iothubId" tf:"iothub_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RangeRule []RangeRuleParameters `json:"rangeRule,omitempty" tf:"range_rule,omitempty"`
}

type RangeRuleObservation struct {
}

type RangeRuleParameters struct {

	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// IOTSecurityDeviceGroupSpec defines the desired state of IOTSecurityDeviceGroup
type IOTSecurityDeviceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTSecurityDeviceGroupParameters `json:"forProvider"`
}

// IOTSecurityDeviceGroupStatus defines the observed state of IOTSecurityDeviceGroup.
type IOTSecurityDeviceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTSecurityDeviceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecurityDeviceGroup is the Schema for the IOTSecurityDeviceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IOTSecurityDeviceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTSecurityDeviceGroupSpec   `json:"spec"`
	Status            IOTSecurityDeviceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecurityDeviceGroupList contains a list of IOTSecurityDeviceGroups
type IOTSecurityDeviceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTSecurityDeviceGroup `json:"items"`
}

// Repository type metadata.
var (
	IOTSecurityDeviceGroup_Kind             = "IOTSecurityDeviceGroup"
	IOTSecurityDeviceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTSecurityDeviceGroup_Kind}.String()
	IOTSecurityDeviceGroup_KindAPIVersion   = IOTSecurityDeviceGroup_Kind + "." + CRDGroupVersion.String()
	IOTSecurityDeviceGroup_GroupVersionKind = CRDGroupVersion.WithKind(IOTSecurityDeviceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTSecurityDeviceGroup{}, &IOTSecurityDeviceGroupList{})
}
