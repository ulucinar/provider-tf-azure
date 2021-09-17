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

type EventhubAuthorizationRuleObservation struct {
	PrimaryConnectionString string `json:"primaryConnectionString,omitempty" tf:"primary_connection_string"`

	PrimaryConnectionStringAlias string `json:"primaryConnectionStringAlias,omitempty" tf:"primary_connection_string_alias"`

	PrimaryKey string `json:"primaryKey,omitempty" tf:"primary_key"`

	SecondaryConnectionString string `json:"secondaryConnectionString,omitempty" tf:"secondary_connection_string"`

	SecondaryConnectionStringAlias string `json:"secondaryConnectionStringAlias,omitempty" tf:"secondary_connection_string_alias"`

	SecondaryKey string `json:"secondaryKey,omitempty" tf:"secondary_key"`
}

type EventhubAuthorizationRuleParameters struct {

	// +kubebuilder:validation:Required
	EventhubName string `json:"eventhubName" tf:"eventhub_name"`

	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen"`

	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NamespaceName string `json:"namespaceName" tf:"namespace_name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send"`
}

// EventhubAuthorizationRuleSpec defines the desired state of EventhubAuthorizationRule
type EventhubAuthorizationRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EventhubAuthorizationRuleParameters `json:"forProvider"`
}

// EventhubAuthorizationRuleStatus defines the observed state of EventhubAuthorizationRule.
type EventhubAuthorizationRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EventhubAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubAuthorizationRule is the Schema for the EventhubAuthorizationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventhubAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubAuthorizationRuleSpec   `json:"spec"`
	Status            EventhubAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubAuthorizationRuleList contains a list of EventhubAuthorizationRules
type EventhubAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	EventhubAuthorizationRuleKind             = "EventhubAuthorizationRule"
	EventhubAuthorizationRuleGroupKind        = schema.GroupKind{Group: Group, Kind: EventhubAuthorizationRuleKind}.String()
	EventhubAuthorizationRuleKindAPIVersion   = EventhubAuthorizationRuleKind + "." + GroupVersion.String()
	EventhubAuthorizationRuleGroupVersionKind = GroupVersion.WithKind(EventhubAuthorizationRuleKind)
)

func init() {
	SchemeBuilder.Register(&EventhubAuthorizationRule{}, &EventhubAuthorizationRuleList{})
}
