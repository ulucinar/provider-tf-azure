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

type InterfaceBackendAddressPoolAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InterfaceBackendAddressPoolAssociationParameters struct {

	// +kubebuilder:validation:Required
	BackendAddressPoolID *string `json:"backendAddressPoolId" tf:"backend_address_pool_id,omitempty"`

	// +kubebuilder:validation:Required
	IPConfigurationName *string `json:"ipConfigurationName" tf:"ip_configuration_name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterfaceID *string `json:"networkInterfaceId" tf:"network_interface_id,omitempty"`
}

// InterfaceBackendAddressPoolAssociationSpec defines the desired state of InterfaceBackendAddressPoolAssociation
type InterfaceBackendAddressPoolAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InterfaceBackendAddressPoolAssociationParameters `json:"forProvider"`
}

// InterfaceBackendAddressPoolAssociationStatus defines the observed state of InterfaceBackendAddressPoolAssociation.
type InterfaceBackendAddressPoolAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InterfaceBackendAddressPoolAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceBackendAddressPoolAssociation is the Schema for the InterfaceBackendAddressPoolAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type InterfaceBackendAddressPoolAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterfaceBackendAddressPoolAssociationSpec   `json:"spec"`
	Status            InterfaceBackendAddressPoolAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceBackendAddressPoolAssociationList contains a list of InterfaceBackendAddressPoolAssociations
type InterfaceBackendAddressPoolAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InterfaceBackendAddressPoolAssociation `json:"items"`
}

// Repository type metadata.
var (
	InterfaceBackendAddressPoolAssociation_Kind             = "InterfaceBackendAddressPoolAssociation"
	InterfaceBackendAddressPoolAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InterfaceBackendAddressPoolAssociation_Kind}.String()
	InterfaceBackendAddressPoolAssociation_KindAPIVersion   = InterfaceBackendAddressPoolAssociation_Kind + "." + CRDGroupVersion.String()
	InterfaceBackendAddressPoolAssociation_GroupVersionKind = CRDGroupVersion.WithKind(InterfaceBackendAddressPoolAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&InterfaceBackendAddressPoolAssociation{}, &InterfaceBackendAddressPoolAssociationList{})
}
