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

type DevTestWindowsVirtualMachineGalleryImageReferenceObservation struct {
}

type DevTestWindowsVirtualMachineGalleryImageReferenceParameters struct {

	// +kubebuilder:validation:Required
	Offer string `json:"offer" tf:"offer"`

	// +kubebuilder:validation:Required
	Publisher string `json:"publisher" tf:"publisher"`

	// +kubebuilder:validation:Required
	Sku string `json:"sku" tf:"sku"`

	// +kubebuilder:validation:Required
	Version string `json:"version" tf:"version"`
}

type DevTestWindowsVirtualMachineInboundNatRuleObservation struct {
	FrontendPort int64 `json:"frontendPort,omitempty" tf:"frontend_port"`
}

type DevTestWindowsVirtualMachineInboundNatRuleParameters struct {

	// +kubebuilder:validation:Required
	BackendPort int64 `json:"backendPort" tf:"backend_port"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`
}

type DevTestWindowsVirtualMachineObservation struct {
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn"`

	UniqueIdentifier string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier"`
}

type DevTestWindowsVirtualMachineParameters struct {

	// +kubebuilder:validation:Optional
	AllowClaim *bool `json:"allowClaim,omitempty" tf:"allow_claim"`

	// +kubebuilder:validation:Optional
	DisallowPublicIPAddress *bool `json:"disallowPublicIpAddress,omitempty" tf:"disallow_public_ip_address"`

	// +kubebuilder:validation:Required
	GalleryImageReference []DevTestWindowsVirtualMachineGalleryImageReferenceParameters `json:"galleryImageReference" tf:"gallery_image_reference"`

	// +kubebuilder:validation:Optional
	InboundNatRule []DevTestWindowsVirtualMachineInboundNatRuleParameters `json:"inboundNatRule,omitempty" tf:"inbound_nat_rule"`

	// +kubebuilder:validation:Required
	LabName string `json:"labName" tf:"lab_name"`

	// +kubebuilder:validation:Required
	LabSubnetName string `json:"labSubnetName" tf:"lab_subnet_name"`

	// +kubebuilder:validation:Required
	LabVirtualNetworkID string `json:"labVirtualNetworkId" tf:"lab_virtual_network_id"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Size string `json:"size" tf:"size"`

	// +kubebuilder:validation:Required
	StorageType string `json:"storageType" tf:"storage_type"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

// DevTestWindowsVirtualMachineSpec defines the desired state of DevTestWindowsVirtualMachine
type DevTestWindowsVirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DevTestWindowsVirtualMachineParameters `json:"forProvider"`
}

// DevTestWindowsVirtualMachineStatus defines the observed state of DevTestWindowsVirtualMachine.
type DevTestWindowsVirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DevTestWindowsVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DevTestWindowsVirtualMachine is the Schema for the DevTestWindowsVirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DevTestWindowsVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestWindowsVirtualMachineSpec   `json:"spec"`
	Status            DevTestWindowsVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevTestWindowsVirtualMachineList contains a list of DevTestWindowsVirtualMachines
type DevTestWindowsVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevTestWindowsVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	DevTestWindowsVirtualMachineKind             = "DevTestWindowsVirtualMachine"
	DevTestWindowsVirtualMachineGroupKind        = schema.GroupKind{Group: Group, Kind: DevTestWindowsVirtualMachineKind}.String()
	DevTestWindowsVirtualMachineKindAPIVersion   = DevTestWindowsVirtualMachineKind + "." + GroupVersion.String()
	DevTestWindowsVirtualMachineGroupVersionKind = GroupVersion.WithKind(DevTestWindowsVirtualMachineKind)
)

func init() {
	SchemeBuilder.Register(&DevTestWindowsVirtualMachine{}, &DevTestWindowsVirtualMachineList{})
}
