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

type EventhubClusterObservation struct {
}

type EventhubClusterParameters struct {

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	SkuName string `json:"skuName" tf:"sku_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// EventhubClusterSpec defines the desired state of EventhubCluster
type EventhubClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EventhubClusterParameters `json:"forProvider"`
}

// EventhubClusterStatus defines the observed state of EventhubCluster.
type EventhubClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EventhubClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubCluster is the Schema for the EventhubClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventhubCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubClusterSpec   `json:"spec"`
	Status            EventhubClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubClusterList contains a list of EventhubClusters
type EventhubClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubCluster `json:"items"`
}

// Repository type metadata.
var (
	EventhubClusterKind             = "EventhubCluster"
	EventhubClusterGroupKind        = schema.GroupKind{Group: Group, Kind: EventhubClusterKind}.String()
	EventhubClusterKindAPIVersion   = EventhubClusterKind + "." + GroupVersion.String()
	EventhubClusterGroupVersionKind = GroupVersion.WithKind(EventhubClusterKind)
)

func init() {
	SchemeBuilder.Register(&EventhubCluster{}, &EventhubClusterList{})
}
