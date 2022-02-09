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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SiteRecoveryProtectionContainerMappingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SiteRecoveryProtectionContainerMappingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryFabricName *string `json:"recoveryFabricName" tf:"recovery_fabric_name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryReplicationPolicyID *string `json:"recoveryReplicationPolicyId" tf:"recovery_replication_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoverySourceProtectionContainerName *string `json:"recoverySourceProtectionContainerName" tf:"recovery_source_protection_container_name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryTargetProtectionContainerID *string `json:"recoveryTargetProtectionContainerId" tf:"recovery_target_protection_container_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// SiteRecoveryProtectionContainerMappingSpec defines the desired state of SiteRecoveryProtectionContainerMapping
type SiteRecoveryProtectionContainerMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryProtectionContainerMappingParameters `json:"forProvider"`
}

// SiteRecoveryProtectionContainerMappingStatus defines the observed state of SiteRecoveryProtectionContainerMapping.
type SiteRecoveryProtectionContainerMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryProtectionContainerMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainerMapping is the Schema for the SiteRecoveryProtectionContainerMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SiteRecoveryProtectionContainerMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryProtectionContainerMappingSpec   `json:"spec"`
	Status            SiteRecoveryProtectionContainerMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainerMappingList contains a list of SiteRecoveryProtectionContainerMappings
type SiteRecoveryProtectionContainerMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryProtectionContainerMapping `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryProtectionContainerMapping_Kind             = "SiteRecoveryProtectionContainerMapping"
	SiteRecoveryProtectionContainerMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryProtectionContainerMapping_Kind}.String()
	SiteRecoveryProtectionContainerMapping_KindAPIVersion   = SiteRecoveryProtectionContainerMapping_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryProtectionContainerMapping_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryProtectionContainerMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryProtectionContainerMapping{}, &SiteRecoveryProtectionContainerMappingList{})
}
