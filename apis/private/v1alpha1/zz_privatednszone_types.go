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

type PrivateDnsZoneObservation struct {
	MaxNumberOfRecordSets int64 `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets"`

	MaxNumberOfVirtualNetworkLinks int64 `json:"maxNumberOfVirtualNetworkLinks,omitempty" tf:"max_number_of_virtual_network_links"`

	MaxNumberOfVirtualNetworkLinksWithRegistration int64 `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty" tf:"max_number_of_virtual_network_links_with_registration"`

	NumberOfRecordSets int64 `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets"`
}

type PrivateDnsZoneParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SoaRecord []SoaRecordParameters `json:"soaRecord,omitempty" tf:"soa_record"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type SoaRecordObservation struct {
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn"`

	HostName string `json:"hostName,omitempty" tf:"host_name"`

	SerialNumber int64 `json:"serialNumber,omitempty" tf:"serial_number"`
}

type SoaRecordParameters struct {

	// +kubebuilder:validation:Required
	Email string `json:"email" tf:"email"`

	// +kubebuilder:validation:Optional
	ExpireTime *int64 `json:"expireTime,omitempty" tf:"expire_time"`

	// +kubebuilder:validation:Optional
	MinimumTTL *int64 `json:"minimumTtl,omitempty" tf:"minimum_ttl"`

	// +kubebuilder:validation:Optional
	RefreshTime *int64 `json:"refreshTime,omitempty" tf:"refresh_time"`

	// +kubebuilder:validation:Optional
	RetryTime *int64 `json:"retryTime,omitempty" tf:"retry_time"`

	// +kubebuilder:validation:Optional
	TTL *int64 `json:"ttl,omitempty" tf:"ttl"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// PrivateDnsZoneSpec defines the desired state of PrivateDnsZone
type PrivateDnsZoneSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PrivateDnsZoneParameters `json:"forProvider"`
}

// PrivateDnsZoneStatus defines the observed state of PrivateDnsZone.
type PrivateDnsZoneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PrivateDnsZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDnsZone is the Schema for the PrivateDnsZones API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDnsZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZoneSpec   `json:"spec"`
	Status            PrivateDnsZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDnsZoneList contains a list of PrivateDnsZones
type PrivateDnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZone `json:"items"`
}

// Repository type metadata.
var (
	PrivateDnsZoneKind             = "PrivateDnsZone"
	PrivateDnsZoneGroupKind        = schema.GroupKind{Group: Group, Kind: PrivateDnsZoneKind}.String()
	PrivateDnsZoneKindAPIVersion   = PrivateDnsZoneKind + "." + GroupVersion.String()
	PrivateDnsZoneGroupVersionKind = GroupVersion.WithKind(PrivateDnsZoneKind)
)

func init() {
	SchemeBuilder.Register(&PrivateDnsZone{}, &PrivateDnsZoneList{})
}
