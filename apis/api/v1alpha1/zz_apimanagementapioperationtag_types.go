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

type ApiManagementApiOperationTagObservation struct {
}

type ApiManagementApiOperationTagParameters struct {

	// +kubebuilder:validation:Required
	APIOperationID string `json:"apiOperationId" tf:"api_operation_id"`

	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName" tf:"display_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

// ApiManagementApiOperationTagSpec defines the desired state of ApiManagementApiOperationTag
type ApiManagementApiOperationTagSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementApiOperationTagParameters `json:"forProvider"`
}

// ApiManagementApiOperationTagStatus defines the observed state of ApiManagementApiOperationTag.
type ApiManagementApiOperationTagStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementApiOperationTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiOperationTag is the Schema for the ApiManagementApiOperationTags API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementApiOperationTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiOperationTagSpec   `json:"spec"`
	Status            ApiManagementApiOperationTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiOperationTagList contains a list of ApiManagementApiOperationTags
type ApiManagementApiOperationTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementApiOperationTag `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementApiOperationTagKind             = "ApiManagementApiOperationTag"
	ApiManagementApiOperationTagGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementApiOperationTagKind}.String()
	ApiManagementApiOperationTagKindAPIVersion   = ApiManagementApiOperationTagKind + "." + GroupVersion.String()
	ApiManagementApiOperationTagGroupVersionKind = GroupVersion.WithKind(ApiManagementApiOperationTagKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementApiOperationTag{}, &ApiManagementApiOperationTagList{})
}
