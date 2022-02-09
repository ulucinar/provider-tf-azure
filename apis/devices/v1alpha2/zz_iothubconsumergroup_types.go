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

type IOTHubConsumerGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubConsumerGroupParameters struct {

	// +kubebuilder:validation:Required
	EventHubEndpointName *string `json:"eventhubEndpointName" tf:"eventhub_endpoint_name,omitempty"`

	// +crossplane:generate:reference:type=IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iotHubNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iotHubNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IOTHubConsumerGroupSpec defines the desired state of IOTHubConsumerGroup
type IOTHubConsumerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubConsumerGroupParameters `json:"forProvider"`
}

// IOTHubConsumerGroupStatus defines the observed state of IOTHubConsumerGroup.
type IOTHubConsumerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubConsumerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubConsumerGroup is the Schema for the IOTHubConsumerGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IOTHubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubConsumerGroupSpec   `json:"spec"`
	Status            IOTHubConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubConsumerGroupList contains a list of IOTHubConsumerGroups
type IOTHubConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubConsumerGroup `json:"items"`
}

// Repository type metadata.
var (
	IOTHubConsumerGroup_Kind             = "IOTHubConsumerGroup"
	IOTHubConsumerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubConsumerGroup_Kind}.String()
	IOTHubConsumerGroup_KindAPIVersion   = IOTHubConsumerGroup_Kind + "." + CRDGroupVersion.String()
	IOTHubConsumerGroup_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubConsumerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubConsumerGroup{}, &IOTHubConsumerGroupList{})
}
