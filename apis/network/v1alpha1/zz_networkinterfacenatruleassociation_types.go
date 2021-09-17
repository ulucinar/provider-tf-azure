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

type NetworkInterfaceNatRuleAssociationObservation struct {
}

type NetworkInterfaceNatRuleAssociationParameters struct {

	// +kubebuilder:validation:Required
	IPConfigurationName string `json:"ipConfigurationName" tf:"ip_configuration_name"`

	// +kubebuilder:validation:Required
	NatRuleID string `json:"natRuleId" tf:"nat_rule_id"`

	// +kubebuilder:validation:Required
	NetworkInterfaceID string `json:"networkInterfaceId" tf:"network_interface_id"`
}

// NetworkInterfaceNatRuleAssociationSpec defines the desired state of NetworkInterfaceNatRuleAssociation
type NetworkInterfaceNatRuleAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NetworkInterfaceNatRuleAssociationParameters `json:"forProvider"`
}

// NetworkInterfaceNatRuleAssociationStatus defines the observed state of NetworkInterfaceNatRuleAssociation.
type NetworkInterfaceNatRuleAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NetworkInterfaceNatRuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceNatRuleAssociation is the Schema for the NetworkInterfaceNatRuleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceNatRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceNatRuleAssociationSpec   `json:"spec"`
	Status            NetworkInterfaceNatRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceNatRuleAssociationList contains a list of NetworkInterfaceNatRuleAssociations
type NetworkInterfaceNatRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceNatRuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceNatRuleAssociationKind             = "NetworkInterfaceNatRuleAssociation"
	NetworkInterfaceNatRuleAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkInterfaceNatRuleAssociationKind}.String()
	NetworkInterfaceNatRuleAssociationKindAPIVersion   = NetworkInterfaceNatRuleAssociationKind + "." + GroupVersion.String()
	NetworkInterfaceNatRuleAssociationGroupVersionKind = GroupVersion.WithKind(NetworkInterfaceNatRuleAssociationKind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceNatRuleAssociation{}, &NetworkInterfaceNatRuleAssociationList{})
}
