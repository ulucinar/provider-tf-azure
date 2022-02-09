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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DDOSProtectionPlanObservation struct {
}

type DDOSProtectionPlanParameters struct {

	// +kubebuilder:validation:Required
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type VirtualNetworkObservation struct {
	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualNetworkParameters struct {

	// +kubebuilder:validation:Required
	AddressSpace []*string `json:"addressSpace" tf:"address_space,omitempty"`

	// +kubebuilder:validation:Optional
	BGPCommunity *string `json:"bgpCommunity,omitempty" tf:"bgp_community,omitempty"`

	// +kubebuilder:validation:Optional
	DDOSProtectionPlan []DDOSProtectionPlanParameters `json:"ddosProtectionPlan,omitempty" tf:"ddos_protection_plan,omitempty"`

	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Subnet []VirtualNetworkSubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VMProtectionEnabled *bool `json:"vmProtectionEnabled,omitempty" tf:"vm_protection_enabled,omitempty"`
}

type VirtualNetworkSubnetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualNetworkSubnetParameters struct {

	// +kubebuilder:validation:Required
	AddressPrefix *string `json:"addressPrefix" tf:"address_prefix,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`
}

// VirtualNetworkSpec defines the desired state of VirtualNetwork
type VirtualNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkParameters `json:"forProvider"`
}

// VirtualNetworkStatus defines the observed state of VirtualNetwork.
type VirtualNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetwork is the Schema for the VirtualNetworks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkSpec   `json:"spec"`
	Status            VirtualNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkList contains a list of VirtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetwork_Kind             = "VirtualNetwork"
	VirtualNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetwork_Kind}.String()
	VirtualNetwork_KindAPIVersion   = VirtualNetwork_Kind + "." + CRDGroupVersion.String()
	VirtualNetwork_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}
