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

type ApiManagementIdentityProviderTwitterObservation struct {
}

type ApiManagementIdentityProviderTwitterParameters struct {

	// +kubebuilder:validation:Required
	APIKey string `json:"apiKey" tf:"api_key"`

	// +kubebuilder:validation:Required
	APIManagementName string `json:"apiManagementName" tf:"api_management_name"`

	// +kubebuilder:validation:Required
	APISecretKey string `json:"apiSecretKey" tf:"api_secret_key"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

// ApiManagementIdentityProviderTwitterSpec defines the desired state of ApiManagementIdentityProviderTwitter
type ApiManagementIdentityProviderTwitterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementIdentityProviderTwitterParameters `json:"forProvider"`
}

// ApiManagementIdentityProviderTwitterStatus defines the observed state of ApiManagementIdentityProviderTwitter.
type ApiManagementIdentityProviderTwitterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementIdentityProviderTwitterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementIdentityProviderTwitter is the Schema for the ApiManagementIdentityProviderTwitters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementIdentityProviderTwitter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementIdentityProviderTwitterSpec   `json:"spec"`
	Status            ApiManagementIdentityProviderTwitterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementIdentityProviderTwitterList contains a list of ApiManagementIdentityProviderTwitters
type ApiManagementIdentityProviderTwitterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementIdentityProviderTwitter `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementIdentityProviderTwitterKind             = "ApiManagementIdentityProviderTwitter"
	ApiManagementIdentityProviderTwitterGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementIdentityProviderTwitterKind}.String()
	ApiManagementIdentityProviderTwitterKindAPIVersion   = ApiManagementIdentityProviderTwitterKind + "." + GroupVersion.String()
	ApiManagementIdentityProviderTwitterGroupVersionKind = GroupVersion.WithKind(ApiManagementIdentityProviderTwitterKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementIdentityProviderTwitter{}, &ApiManagementIdentityProviderTwitterList{})
}
