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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type IothubDpsSharedAccessPolicyObservation struct {
	PrimaryConnectionString string `json:"primaryConnectionString,omitempty" tf:"primary_connection_string"`

	PrimaryKey string `json:"primaryKey,omitempty" tf:"primary_key"`

	SecondaryConnectionString string `json:"secondaryConnectionString,omitempty" tf:"secondary_connection_string"`

	SecondaryKey string `json:"secondaryKey,omitempty" tf:"secondary_key"`
}

type IothubDpsSharedAccessPolicyParameters struct {

	// +kubebuilder:validation:Optional
	EnrollmentRead *bool `json:"enrollmentRead,omitempty" tf:"enrollment_read"`

	// +kubebuilder:validation:Optional
	EnrollmentWrite *bool `json:"enrollmentWrite,omitempty" tf:"enrollment_write"`

	// +kubebuilder:validation:Required
	IothubDpsName string `json:"iothubDpsName" tf:"iothub_dps_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	RegistrationRead *bool `json:"registrationRead,omitempty" tf:"registration_read"`

	// +kubebuilder:validation:Optional
	RegistrationWrite *bool `json:"registrationWrite,omitempty" tf:"registration_write"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ServiceConfig *bool `json:"serviceConfig,omitempty" tf:"service_config"`
}

// IothubDpsSharedAccessPolicySpec defines the desired state of IothubDpsSharedAccessPolicy
type IothubDpsSharedAccessPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IothubDpsSharedAccessPolicyParameters `json:"forProvider"`
}

// IothubDpsSharedAccessPolicyStatus defines the observed state of IothubDpsSharedAccessPolicy.
type IothubDpsSharedAccessPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IothubDpsSharedAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IothubDpsSharedAccessPolicy is the Schema for the IothubDpsSharedAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IothubDpsSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubDpsSharedAccessPolicySpec   `json:"spec"`
	Status            IothubDpsSharedAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubDpsSharedAccessPolicyList contains a list of IothubDpsSharedAccessPolicys
type IothubDpsSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IothubDpsSharedAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	IothubDpsSharedAccessPolicyKind             = "IothubDpsSharedAccessPolicy"
	IothubDpsSharedAccessPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: IothubDpsSharedAccessPolicyKind}.String()
	IothubDpsSharedAccessPolicyKindAPIVersion   = IothubDpsSharedAccessPolicyKind + "." + GroupVersion.String()
	IothubDpsSharedAccessPolicyGroupVersionKind = GroupVersion.WithKind(IothubDpsSharedAccessPolicyKind)
)

func init() {
	SchemeBuilder.Register(&IothubDpsSharedAccessPolicy{}, &IothubDpsSharedAccessPolicyList{})
}
