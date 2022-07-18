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

type MSSQLDatabaseExtendedAuditingPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MSSQLDatabaseExtendedAuditingPolicyParameters struct {

	// +kubebuilder:validation:Required
	DatabaseID *string `json:"databaseId" tf:"database_id,omitempty"`

	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// MSSQLDatabaseExtendedAuditingPolicySpec defines the desired state of MSSQLDatabaseExtendedAuditingPolicy
type MSSQLDatabaseExtendedAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLDatabaseExtendedAuditingPolicyParameters `json:"forProvider"`
}

// MSSQLDatabaseExtendedAuditingPolicyStatus defines the observed state of MSSQLDatabaseExtendedAuditingPolicy.
type MSSQLDatabaseExtendedAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLDatabaseExtendedAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLDatabaseExtendedAuditingPolicy is the Schema for the MSSQLDatabaseExtendedAuditingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MSSQLDatabaseExtendedAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLDatabaseExtendedAuditingPolicySpec   `json:"spec"`
	Status            MSSQLDatabaseExtendedAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLDatabaseExtendedAuditingPolicyList contains a list of MSSQLDatabaseExtendedAuditingPolicys
type MSSQLDatabaseExtendedAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLDatabaseExtendedAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	MSSQLDatabaseExtendedAuditingPolicy_Kind             = "MSSQLDatabaseExtendedAuditingPolicy"
	MSSQLDatabaseExtendedAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLDatabaseExtendedAuditingPolicy_Kind}.String()
	MSSQLDatabaseExtendedAuditingPolicy_KindAPIVersion   = MSSQLDatabaseExtendedAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	MSSQLDatabaseExtendedAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLDatabaseExtendedAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLDatabaseExtendedAuditingPolicy{}, &MSSQLDatabaseExtendedAuditingPolicyList{})
}
