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

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string"`

	// +kubebuilder:validation:Required
	ResourceID string `json:"resourceId" tf:"resource_id"`

	// +kubebuilder:validation:Optional
	TriggerURL *string `json:"triggerUrl,omitempty" tf:"trigger_url"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Required
	ExpectedValue string `json:"expectedValue" tf:"expected_value"`

	// +kubebuilder:validation:Required
	Operator string `json:"operator" tf:"operator"`

	// +kubebuilder:validation:Required
	PropertyPath string `json:"propertyPath" tf:"property_path"`

	// +kubebuilder:validation:Required
	PropertyType string `json:"propertyType" tf:"property_type"`
}

type RuleSetObservation struct {
}

type RuleSetParameters struct {

	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule"`
}

type SecurityCenterAutomationObservation struct {
}

type SecurityCenterAutomationParameters struct {

	// +kubebuilder:validation:Required
	Action []ActionParameters `json:"action" tf:"action"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Scopes []string `json:"scopes" tf:"scopes"`

	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	EventSource string `json:"eventSource" tf:"event_source"`

	// +kubebuilder:validation:Optional
	RuleSet []RuleSetParameters `json:"ruleSet,omitempty" tf:"rule_set"`
}

// SecurityCenterAutomationSpec defines the desired state of SecurityCenterAutomation
type SecurityCenterAutomationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecurityCenterAutomationParameters `json:"forProvider"`
}

// SecurityCenterAutomationStatus defines the observed state of SecurityCenterAutomation.
type SecurityCenterAutomationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecurityCenterAutomationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAutomation is the Schema for the SecurityCenterAutomations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterAutomation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterAutomationSpec   `json:"spec"`
	Status            SecurityCenterAutomationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAutomationList contains a list of SecurityCenterAutomations
type SecurityCenterAutomationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterAutomation `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterAutomationKind             = "SecurityCenterAutomation"
	SecurityCenterAutomationGroupKind        = schema.GroupKind{Group: Group, Kind: SecurityCenterAutomationKind}.String()
	SecurityCenterAutomationKindAPIVersion   = SecurityCenterAutomationKind + "." + GroupVersion.String()
	SecurityCenterAutomationGroupVersionKind = GroupVersion.WithKind(SecurityCenterAutomationKind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterAutomation{}, &SecurityCenterAutomationList{})
}
