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

type FirewallObservation struct {
}

type FirewallParameters struct {

	// +kubebuilder:validation:Optional
	DNSServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`

	// +kubebuilder:validation:Optional
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id"`

	// +kubebuilder:validation:Optional
	IPConfiguration []IPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Optional
	ManagementIPConfiguration []ManagementIPConfigurationParameters `json:"managementIpConfiguration,omitempty" tf:"management_ip_configuration"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PrivateIPRanges []string `json:"privateIpRanges,omitempty" tf:"private_ip_ranges"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name"`

	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode"`

	// +kubebuilder:validation:Optional
	VirtualHub []VirtualHubParameters `json:"virtualHub,omitempty" tf:"virtual_hub"`

	// +kubebuilder:validation:Optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type IPConfigurationObservation struct {
	PrivateIPAddress string `json:"privateIpAddress,omitempty" tf:"private_ip_address"`
}

type IPConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	PublicIPAddressID string `json:"publicIpAddressId" tf:"public_ip_address_id"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`
}

type ManagementIPConfigurationObservation struct {
	PrivateIPAddress string `json:"privateIpAddress,omitempty" tf:"private_ip_address"`
}

type ManagementIPConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	PublicIPAddressID string `json:"publicIpAddressId" tf:"public_ip_address_id"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

type VirtualHubObservation struct {
	PrivateIPAddress string `json:"privateIpAddress,omitempty" tf:"private_ip_address"`

	PublicIPAddresses []string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses"`
}

type VirtualHubParameters struct {

	// +kubebuilder:validation:Optional
	PublicIPCount *int64 `json:"publicIpCount,omitempty" tf:"public_ip_count"`

	// +kubebuilder:validation:Required
	VirtualHubID string `json:"virtualHubId" tf:"virtual_hub_id"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FirewallParameters `json:"forProvider"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec"`
	Status            FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	FirewallKind             = "Firewall"
	FirewallGroupKind        = schema.GroupKind{Group: Group, Kind: FirewallKind}.String()
	FirewallKindAPIVersion   = FirewallKind + "." + GroupVersion.String()
	FirewallGroupVersionKind = GroupVersion.WithKind(FirewallKind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
