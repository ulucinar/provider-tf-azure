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

type DiskEncryptionSetObservation struct {
}

type DiskEncryptionSetParameters struct {

	// +kubebuilder:validation:Required
	Identity []IdentityParameters `json:"identity" tf:"identity"`

	// +kubebuilder:validation:Required
	KeyVaultKeyID string `json:"keyVaultKeyId" tf:"key_vault_key_id"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type IdentityObservation struct {
	PrincipalID string `json:"principalId,omitempty" tf:"principal_id"`

	TenantID string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

// DiskEncryptionSetSpec defines the desired state of DiskEncryptionSet
type DiskEncryptionSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DiskEncryptionSetParameters `json:"forProvider"`
}

// DiskEncryptionSetStatus defines the observed state of DiskEncryptionSet.
type DiskEncryptionSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DiskEncryptionSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskEncryptionSet is the Schema for the DiskEncryptionSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DiskEncryptionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskEncryptionSetSpec   `json:"spec"`
	Status            DiskEncryptionSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskEncryptionSetList contains a list of DiskEncryptionSets
type DiskEncryptionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskEncryptionSet `json:"items"`
}

// Repository type metadata.
var (
	DiskEncryptionSetKind             = "DiskEncryptionSet"
	DiskEncryptionSetGroupKind        = schema.GroupKind{Group: Group, Kind: DiskEncryptionSetKind}.String()
	DiskEncryptionSetKindAPIVersion   = DiskEncryptionSetKind + "." + GroupVersion.String()
	DiskEncryptionSetGroupVersionKind = GroupVersion.WithKind(DiskEncryptionSetKind)
)

func init() {
	SchemeBuilder.Register(&DiskEncryptionSet{}, &DiskEncryptionSetList{})
}
