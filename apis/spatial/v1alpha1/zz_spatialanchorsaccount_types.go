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

type SpatialAnchorsAccountObservation struct {
	AccountDomain string `json:"accountDomain,omitempty" tf:"account_domain"`

	AccountID string `json:"accountId,omitempty" tf:"account_id"`
}

type SpatialAnchorsAccountParameters struct {

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// SpatialAnchorsAccountSpec defines the desired state of SpatialAnchorsAccount
type SpatialAnchorsAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SpatialAnchorsAccountParameters `json:"forProvider"`
}

// SpatialAnchorsAccountStatus defines the observed state of SpatialAnchorsAccount.
type SpatialAnchorsAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SpatialAnchorsAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpatialAnchorsAccount is the Schema for the SpatialAnchorsAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpatialAnchorsAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpatialAnchorsAccountSpec   `json:"spec"`
	Status            SpatialAnchorsAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpatialAnchorsAccountList contains a list of SpatialAnchorsAccounts
type SpatialAnchorsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpatialAnchorsAccount `json:"items"`
}

// Repository type metadata.
var (
	SpatialAnchorsAccountKind             = "SpatialAnchorsAccount"
	SpatialAnchorsAccountGroupKind        = schema.GroupKind{Group: Group, Kind: SpatialAnchorsAccountKind}.String()
	SpatialAnchorsAccountKindAPIVersion   = SpatialAnchorsAccountKind + "." + GroupVersion.String()
	SpatialAnchorsAccountGroupVersionKind = GroupVersion.WithKind(SpatialAnchorsAccountKind)
)

func init() {
	SchemeBuilder.Register(&SpatialAnchorsAccount{}, &SpatialAnchorsAccountList{})
}
