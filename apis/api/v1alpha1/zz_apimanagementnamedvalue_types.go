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

type ApiManagementNamedValueObservation struct {
}

type ApiManagementNamedValueParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName string `json:"apiManagementName" tf:"api_management_name"`

	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName" tf:"display_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Secret *bool `json:"secret,omitempty" tf:"secret"`

	// +kubebuilder:validation:Optional
	Tags []string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value"`

	// +kubebuilder:validation:Optional
	ValueFromKeyVault []ValueFromKeyVaultParameters `json:"valueFromKeyVault,omitempty" tf:"value_from_key_vault"`
}

type ValueFromKeyVaultObservation struct {
}

type ValueFromKeyVaultParameters struct {

	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id"`

	// +kubebuilder:validation:Required
	SecretID string `json:"secretId" tf:"secret_id"`
}

// ApiManagementNamedValueSpec defines the desired state of ApiManagementNamedValue
type ApiManagementNamedValueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementNamedValueParameters `json:"forProvider"`
}

// ApiManagementNamedValueStatus defines the observed state of ApiManagementNamedValue.
type ApiManagementNamedValueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementNamedValueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementNamedValue is the Schema for the ApiManagementNamedValues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementNamedValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementNamedValueSpec   `json:"spec"`
	Status            ApiManagementNamedValueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementNamedValueList contains a list of ApiManagementNamedValues
type ApiManagementNamedValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementNamedValue `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementNamedValueKind             = "ApiManagementNamedValue"
	ApiManagementNamedValueGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementNamedValueKind}.String()
	ApiManagementNamedValueKindAPIVersion   = ApiManagementNamedValueKind + "." + GroupVersion.String()
	ApiManagementNamedValueGroupVersionKind = GroupVersion.WithKind(ApiManagementNamedValueKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementNamedValue{}, &ApiManagementNamedValueList{})
}
