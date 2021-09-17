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

type PostgresqlVirtualNetworkRuleObservation struct {
}

type PostgresqlVirtualNetworkRuleParameters struct {

	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ServerName string `json:"serverName" tf:"server_name"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

// PostgresqlVirtualNetworkRuleSpec defines the desired state of PostgresqlVirtualNetworkRule
type PostgresqlVirtualNetworkRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostgresqlVirtualNetworkRuleParameters `json:"forProvider"`
}

// PostgresqlVirtualNetworkRuleStatus defines the observed state of PostgresqlVirtualNetworkRule.
type PostgresqlVirtualNetworkRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostgresqlVirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlVirtualNetworkRule is the Schema for the PostgresqlVirtualNetworkRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlVirtualNetworkRuleSpec   `json:"spec"`
	Status            PostgresqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlVirtualNetworkRuleList contains a list of PostgresqlVirtualNetworkRules
type PostgresqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlVirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlVirtualNetworkRuleKind             = "PostgresqlVirtualNetworkRule"
	PostgresqlVirtualNetworkRuleGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlVirtualNetworkRuleKind}.String()
	PostgresqlVirtualNetworkRuleKindAPIVersion   = PostgresqlVirtualNetworkRuleKind + "." + GroupVersion.String()
	PostgresqlVirtualNetworkRuleGroupVersionKind = GroupVersion.WithKind(PostgresqlVirtualNetworkRuleKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlVirtualNetworkRule{}, &PostgresqlVirtualNetworkRuleList{})
}
