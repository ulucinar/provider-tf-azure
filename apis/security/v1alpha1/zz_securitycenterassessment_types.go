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

type SecurityCenterAssessmentObservation struct {
}

type SecurityCenterAssessmentParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalData map[string]string `json:"additionalData,omitempty" tf:"additional_data"`

	// +kubebuilder:validation:Required
	AssessmentPolicyID string `json:"assessmentPolicyId" tf:"assessment_policy_id"`

	// +kubebuilder:validation:Required
	Status []StatusParameters `json:"status" tf:"status"`

	// +kubebuilder:validation:Required
	TargetResourceID string `json:"targetResourceId" tf:"target_resource_id"`
}

type StatusObservation struct {
}

type StatusParameters struct {

	// +kubebuilder:validation:Optional
	Cause *string `json:"cause,omitempty" tf:"cause"`

	// +kubebuilder:validation:Required
	Code string `json:"code" tf:"code"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`
}

// SecurityCenterAssessmentSpec defines the desired state of SecurityCenterAssessment
type SecurityCenterAssessmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecurityCenterAssessmentParameters `json:"forProvider"`
}

// SecurityCenterAssessmentStatus defines the observed state of SecurityCenterAssessment.
type SecurityCenterAssessmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecurityCenterAssessmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessment is the Schema for the SecurityCenterAssessments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterAssessment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterAssessmentSpec   `json:"spec"`
	Status            SecurityCenterAssessmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessmentList contains a list of SecurityCenterAssessments
type SecurityCenterAssessmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterAssessment `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterAssessmentKind             = "SecurityCenterAssessment"
	SecurityCenterAssessmentGroupKind        = schema.GroupKind{Group: Group, Kind: SecurityCenterAssessmentKind}.String()
	SecurityCenterAssessmentKindAPIVersion   = SecurityCenterAssessmentKind + "." + GroupVersion.String()
	SecurityCenterAssessmentGroupVersionKind = GroupVersion.WithKind(SecurityCenterAssessmentKind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterAssessment{}, &SecurityCenterAssessmentList{})
}
