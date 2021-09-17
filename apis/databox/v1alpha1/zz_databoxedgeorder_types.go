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

type ContactObservation struct {
}

type ContactParameters struct {

	// +kubebuilder:validation:Required
	CompanyName string `json:"companyName" tf:"company_name"`

	// +kubebuilder:validation:Required
	Emails []string `json:"emails" tf:"emails"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	PhoneNumber string `json:"phoneNumber" tf:"phone_number"`
}

type DataboxEdgeOrderObservation struct {
	Name string `json:"name,omitempty" tf:"name"`

	ReturnTracking []ReturnTrackingObservation `json:"returnTracking,omitempty" tf:"return_tracking"`

	SerialNumber string `json:"serialNumber,omitempty" tf:"serial_number"`

	ShipmentHistory []ShipmentHistoryObservation `json:"shipmentHistory,omitempty" tf:"shipment_history"`

	ShipmentTracking []ShipmentTrackingObservation `json:"shipmentTracking,omitempty" tf:"shipment_tracking"`

	Status []StatusObservation `json:"status,omitempty" tf:"status"`
}

type DataboxEdgeOrderParameters struct {

	// +kubebuilder:validation:Required
	Contact []ContactParameters `json:"contact" tf:"contact"`

	// +kubebuilder:validation:Required
	DeviceName string `json:"deviceName" tf:"device_name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ShipmentAddress []ShipmentAddressParameters `json:"shipmentAddress" tf:"shipment_address"`
}

type ReturnTrackingObservation struct {
	CarrierName string `json:"carrierName,omitempty" tf:"carrier_name"`

	SerialNumber string `json:"serialNumber,omitempty" tf:"serial_number"`

	TrackingID string `json:"trackingId,omitempty" tf:"tracking_id"`

	TrackingURL string `json:"trackingUrl,omitempty" tf:"tracking_url"`
}

type ReturnTrackingParameters struct {
}

type ShipmentAddressObservation struct {
}

type ShipmentAddressParameters struct {

	// +kubebuilder:validation:Required
	Address []string `json:"address" tf:"address"`

	// +kubebuilder:validation:Required
	City string `json:"city" tf:"city"`

	// +kubebuilder:validation:Required
	Country string `json:"country" tf:"country"`

	// +kubebuilder:validation:Required
	PostalCode string `json:"postalCode" tf:"postal_code"`

	// +kubebuilder:validation:Required
	State string `json:"state" tf:"state"`
}

type ShipmentHistoryObservation struct {
	AdditionalDetails map[string]string `json:"additionalDetails,omitempty" tf:"additional_details"`

	Comments string `json:"comments,omitempty" tf:"comments"`

	LastUpdate string `json:"lastUpdate,omitempty" tf:"last_update"`
}

type ShipmentHistoryParameters struct {
}

type ShipmentTrackingObservation struct {
	CarrierName string `json:"carrierName,omitempty" tf:"carrier_name"`

	SerialNumber string `json:"serialNumber,omitempty" tf:"serial_number"`

	TrackingID string `json:"trackingId,omitempty" tf:"tracking_id"`

	TrackingURL string `json:"trackingUrl,omitempty" tf:"tracking_url"`
}

type ShipmentTrackingParameters struct {
}

type StatusObservation struct {
	AdditionalDetails map[string]string `json:"additionalDetails,omitempty" tf:"additional_details"`

	Comments string `json:"comments,omitempty" tf:"comments"`

	Info string `json:"info,omitempty" tf:"info"`

	LastUpdate string `json:"lastUpdate,omitempty" tf:"last_update"`
}

type StatusParameters struct {
}

// DataboxEdgeOrderSpec defines the desired state of DataboxEdgeOrder
type DataboxEdgeOrderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataboxEdgeOrderParameters `json:"forProvider"`
}

// DataboxEdgeOrderStatus defines the observed state of DataboxEdgeOrder.
type DataboxEdgeOrderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataboxEdgeOrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataboxEdgeOrder is the Schema for the DataboxEdgeOrders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataboxEdgeOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataboxEdgeOrderSpec   `json:"spec"`
	Status            DataboxEdgeOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataboxEdgeOrderList contains a list of DataboxEdgeOrders
type DataboxEdgeOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataboxEdgeOrder `json:"items"`
}

// Repository type metadata.
var (
	DataboxEdgeOrderKind             = "DataboxEdgeOrder"
	DataboxEdgeOrderGroupKind        = schema.GroupKind{Group: Group, Kind: DataboxEdgeOrderKind}.String()
	DataboxEdgeOrderKindAPIVersion   = DataboxEdgeOrderKind + "." + GroupVersion.String()
	DataboxEdgeOrderGroupVersionKind = GroupVersion.WithKind(DataboxEdgeOrderKind)
)

func init() {
	SchemeBuilder.Register(&DataboxEdgeOrder{}, &DataboxEdgeOrderList{})
}
