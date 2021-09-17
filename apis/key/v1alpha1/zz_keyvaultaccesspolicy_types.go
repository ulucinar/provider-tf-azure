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

type KeyVaultAccessPolicyObservation struct {
}

type KeyVaultAccessPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id"`

	// +kubebuilder:validation:Optional
	CertificatePermissions []string `json:"certificatePermissions,omitempty" tf:"certificate_permissions"`

	// +kubebuilder:validation:Optional
	KeyPermissions []string `json:"keyPermissions,omitempty" tf:"key_permissions"`

	// +kubebuilder:validation:Required
	KeyVaultID string `json:"keyVaultId" tf:"key_vault_id"`

	// +kubebuilder:validation:Required
	ObjectID string `json:"objectId" tf:"object_id"`

	// +kubebuilder:validation:Optional
	SecretPermissions []string `json:"secretPermissions,omitempty" tf:"secret_permissions"`

	// +kubebuilder:validation:Optional
	StoragePermissions []string `json:"storagePermissions,omitempty" tf:"storage_permissions"`

	// +kubebuilder:validation:Required
	TenantID string `json:"tenantId" tf:"tenant_id"`
}

// KeyVaultAccessPolicySpec defines the desired state of KeyVaultAccessPolicy
type KeyVaultAccessPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KeyVaultAccessPolicyParameters `json:"forProvider"`
}

// KeyVaultAccessPolicyStatus defines the observed state of KeyVaultAccessPolicy.
type KeyVaultAccessPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KeyVaultAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyVaultAccessPolicy is the Schema for the KeyVaultAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KeyVaultAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultAccessPolicySpec   `json:"spec"`
	Status            KeyVaultAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyVaultAccessPolicyList contains a list of KeyVaultAccessPolicys
type KeyVaultAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyVaultAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	KeyVaultAccessPolicyKind             = "KeyVaultAccessPolicy"
	KeyVaultAccessPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: KeyVaultAccessPolicyKind}.String()
	KeyVaultAccessPolicyKindAPIVersion   = KeyVaultAccessPolicyKind + "." + GroupVersion.String()
	KeyVaultAccessPolicyGroupVersionKind = GroupVersion.WithKind(KeyVaultAccessPolicyKind)
)

func init() {
	SchemeBuilder.Register(&KeyVaultAccessPolicy{}, &KeyVaultAccessPolicyList{})
}
