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

type MssqlFirewallRuleObservation struct {
}

type MssqlFirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	EndIPAddress string `json:"endIpAddress" tf:"end_ip_address"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ServerID string `json:"serverId" tf:"server_id"`

	// +kubebuilder:validation:Required
	StartIPAddress string `json:"startIpAddress" tf:"start_ip_address"`
}

// MssqlFirewallRuleSpec defines the desired state of MssqlFirewallRule
type MssqlFirewallRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MssqlFirewallRuleParameters `json:"forProvider"`
}

// MssqlFirewallRuleStatus defines the observed state of MssqlFirewallRule.
type MssqlFirewallRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MssqlFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MssqlFirewallRule is the Schema for the MssqlFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MssqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlFirewallRuleSpec   `json:"spec"`
	Status            MssqlFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MssqlFirewallRuleList contains a list of MssqlFirewallRules
type MssqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MssqlFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	MssqlFirewallRuleKind             = "MssqlFirewallRule"
	MssqlFirewallRuleGroupKind        = schema.GroupKind{Group: Group, Kind: MssqlFirewallRuleKind}.String()
	MssqlFirewallRuleKindAPIVersion   = MssqlFirewallRuleKind + "." + GroupVersion.String()
	MssqlFirewallRuleGroupVersionKind = GroupVersion.WithKind(MssqlFirewallRuleKind)
)

func init() {
	SchemeBuilder.Register(&MssqlFirewallRule{}, &MssqlFirewallRuleList{})
}
