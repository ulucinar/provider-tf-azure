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

type NetworkSecurityRuleObservation struct {
}

type NetworkSecurityRuleParameters struct {

	// +kubebuilder:validation:Required
	Access string `json:"access" tf:"access"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix"`

	// +kubebuilder:validation:Optional
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes"`

	// +kubebuilder:validation:Optional
	DestinationApplicationSecurityGroupIds []string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids"`

	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range"`

	// +kubebuilder:validation:Optional
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges"`

	// +kubebuilder:validation:Required
	Direction string `json:"direction" tf:"direction"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NetworkSecurityGroupName string `json:"networkSecurityGroupName" tf:"network_security_group_name"`

	// +kubebuilder:validation:Required
	Priority int64 `json:"priority" tf:"priority"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix"`

	// +kubebuilder:validation:Optional
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes"`

	// +kubebuilder:validation:Optional
	SourceApplicationSecurityGroupIds []string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids"`

	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range"`

	// +kubebuilder:validation:Optional
	SourcePortRanges []string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges"`
}

// NetworkSecurityRuleSpec defines the desired state of NetworkSecurityRule
type NetworkSecurityRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NetworkSecurityRuleParameters `json:"forProvider"`
}

// NetworkSecurityRuleStatus defines the observed state of NetworkSecurityRule.
type NetworkSecurityRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NetworkSecurityRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityRule is the Schema for the NetworkSecurityRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkSecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityRuleSpec   `json:"spec"`
	Status            NetworkSecurityRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityRuleList contains a list of NetworkSecurityRules
type NetworkSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityRule `json:"items"`
}

// Repository type metadata.
var (
	NetworkSecurityRuleKind             = "NetworkSecurityRule"
	NetworkSecurityRuleGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkSecurityRuleKind}.String()
	NetworkSecurityRuleKindAPIVersion   = NetworkSecurityRuleKind + "." + GroupVersion.String()
	NetworkSecurityRuleGroupVersionKind = GroupVersion.WithKind(NetworkSecurityRuleKind)
)

func init() {
	SchemeBuilder.Register(&NetworkSecurityRule{}, &NetworkSecurityRuleList{})
}
