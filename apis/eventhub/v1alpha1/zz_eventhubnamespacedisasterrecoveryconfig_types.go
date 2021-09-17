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

type EventhubNamespaceDisasterRecoveryConfigObservation struct {
}

type EventhubNamespaceDisasterRecoveryConfigParameters struct {

	// +kubebuilder:validation:Optional
	AlternateName *string `json:"alternateName,omitempty" tf:"alternate_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NamespaceName string `json:"namespaceName" tf:"namespace_name"`

	// +kubebuilder:validation:Required
	PartnerNamespaceID string `json:"partnerNamespaceId" tf:"partner_namespace_id"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

// EventhubNamespaceDisasterRecoveryConfigSpec defines the desired state of EventhubNamespaceDisasterRecoveryConfig
type EventhubNamespaceDisasterRecoveryConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EventhubNamespaceDisasterRecoveryConfigParameters `json:"forProvider"`
}

// EventhubNamespaceDisasterRecoveryConfigStatus defines the observed state of EventhubNamespaceDisasterRecoveryConfig.
type EventhubNamespaceDisasterRecoveryConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EventhubNamespaceDisasterRecoveryConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceDisasterRecoveryConfig is the Schema for the EventhubNamespaceDisasterRecoveryConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventhubNamespaceDisasterRecoveryConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceDisasterRecoveryConfigSpec   `json:"spec"`
	Status            EventhubNamespaceDisasterRecoveryConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceDisasterRecoveryConfigList contains a list of EventhubNamespaceDisasterRecoveryConfigs
type EventhubNamespaceDisasterRecoveryConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubNamespaceDisasterRecoveryConfig `json:"items"`
}

// Repository type metadata.
var (
	EventhubNamespaceDisasterRecoveryConfigKind             = "EventhubNamespaceDisasterRecoveryConfig"
	EventhubNamespaceDisasterRecoveryConfigGroupKind        = schema.GroupKind{Group: Group, Kind: EventhubNamespaceDisasterRecoveryConfigKind}.String()
	EventhubNamespaceDisasterRecoveryConfigKindAPIVersion   = EventhubNamespaceDisasterRecoveryConfigKind + "." + GroupVersion.String()
	EventhubNamespaceDisasterRecoveryConfigGroupVersionKind = GroupVersion.WithKind(EventhubNamespaceDisasterRecoveryConfigKind)
)

func init() {
	SchemeBuilder.Register(&EventhubNamespaceDisasterRecoveryConfig{}, &EventhubNamespaceDisasterRecoveryConfigList{})
}
