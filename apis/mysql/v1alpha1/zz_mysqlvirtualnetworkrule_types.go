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

type MysqlVirtualNetworkRuleObservation struct {
}

type MysqlVirtualNetworkRuleParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ServerName string `json:"serverName" tf:"server_name"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

// MysqlVirtualNetworkRuleSpec defines the desired state of MysqlVirtualNetworkRule
type MysqlVirtualNetworkRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MysqlVirtualNetworkRuleParameters `json:"forProvider"`
}

// MysqlVirtualNetworkRuleStatus defines the observed state of MysqlVirtualNetworkRule.
type MysqlVirtualNetworkRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MysqlVirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MysqlVirtualNetworkRule is the Schema for the MysqlVirtualNetworkRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MysqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MysqlVirtualNetworkRuleSpec   `json:"spec"`
	Status            MysqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MysqlVirtualNetworkRuleList contains a list of MysqlVirtualNetworkRules
type MysqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MysqlVirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	MysqlVirtualNetworkRuleKind             = "MysqlVirtualNetworkRule"
	MysqlVirtualNetworkRuleGroupKind        = schema.GroupKind{Group: Group, Kind: MysqlVirtualNetworkRuleKind}.String()
	MysqlVirtualNetworkRuleKindAPIVersion   = MysqlVirtualNetworkRuleKind + "." + GroupVersion.String()
	MysqlVirtualNetworkRuleGroupVersionKind = GroupVersion.WithKind(MysqlVirtualNetworkRuleKind)
)

func init() {
	SchemeBuilder.Register(&MysqlVirtualNetworkRule{}, &MysqlVirtualNetworkRuleList{})
}
