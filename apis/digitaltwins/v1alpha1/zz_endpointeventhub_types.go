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

type EndpointEventHubObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointEventHubParameters struct {

	// +kubebuilder:validation:Optional
	DeadLetterStorageSecretSecretRef *v1.SecretKeySelector `json:"deadLetterStorageSecretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	DigitalTwinsID *string `json:"digitalTwinsId" tf:"digital_twins_id,omitempty"`

	// +kubebuilder:validation:Required
	EventHubPrimaryConnectionStringSecretRef v1.SecretKeySelector `json:"eventHubPrimaryConnectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	EventHubSecondaryConnectionStringSecretRef v1.SecretKeySelector `json:"eventHubSecondaryConnectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// EndpointEventHubSpec defines the desired state of EndpointEventHub
type EndpointEventHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointEventHubParameters `json:"forProvider"`
}

// EndpointEventHubStatus defines the observed state of EndpointEventHub.
type EndpointEventHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointEventHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointEventHub is the Schema for the EndpointEventHubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EndpointEventHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointEventHubSpec   `json:"spec"`
	Status            EndpointEventHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointEventHubList contains a list of EndpointEventHubs
type EndpointEventHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointEventHub `json:"items"`
}

// Repository type metadata.
var (
	EndpointEventHub_Kind             = "EndpointEventHub"
	EndpointEventHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointEventHub_Kind}.String()
	EndpointEventHub_KindAPIVersion   = EndpointEventHub_Kind + "." + CRDGroupVersion.String()
	EndpointEventHub_GroupVersionKind = CRDGroupVersion.WithKind(EndpointEventHub_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointEventHub{}, &EndpointEventHubList{})
}
