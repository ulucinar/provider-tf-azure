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

type SqlActiveDirectoryAdministratorObservation struct {
}

type SqlActiveDirectoryAdministratorParameters struct {

	// +kubebuilder:validation:Required
	Login string `json:"login" tf:"login"`

	// +kubebuilder:validation:Required
	ObjectID string `json:"objectId" tf:"object_id"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ServerName string `json:"serverName" tf:"server_name"`

	// +kubebuilder:validation:Required
	TenantID string `json:"tenantId" tf:"tenant_id"`
}

// SqlActiveDirectoryAdministratorSpec defines the desired state of SqlActiveDirectoryAdministrator
type SqlActiveDirectoryAdministratorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SqlActiveDirectoryAdministratorParameters `json:"forProvider"`
}

// SqlActiveDirectoryAdministratorStatus defines the observed state of SqlActiveDirectoryAdministrator.
type SqlActiveDirectoryAdministratorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SqlActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlActiveDirectoryAdministrator is the Schema for the SqlActiveDirectoryAdministrators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SqlActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlActiveDirectoryAdministratorSpec   `json:"spec"`
	Status            SqlActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlActiveDirectoryAdministratorList contains a list of SqlActiveDirectoryAdministrators
type SqlActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	SqlActiveDirectoryAdministratorKind             = "SqlActiveDirectoryAdministrator"
	SqlActiveDirectoryAdministratorGroupKind        = schema.GroupKind{Group: Group, Kind: SqlActiveDirectoryAdministratorKind}.String()
	SqlActiveDirectoryAdministratorKindAPIVersion   = SqlActiveDirectoryAdministratorKind + "." + GroupVersion.String()
	SqlActiveDirectoryAdministratorGroupVersionKind = GroupVersion.WithKind(SqlActiveDirectoryAdministratorKind)
)

func init() {
	SchemeBuilder.Register(&SqlActiveDirectoryAdministrator{}, &SqlActiveDirectoryAdministratorList{})
}
