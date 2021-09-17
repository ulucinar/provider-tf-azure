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

type SecurityCenterContactObservation struct {
}

type SecurityCenterContactParameters struct {

	// +kubebuilder:validation:Required
	AlertNotifications bool `json:"alertNotifications" tf:"alert_notifications"`

	// +kubebuilder:validation:Required
	AlertsToAdmins bool `json:"alertsToAdmins" tf:"alerts_to_admins"`

	// +kubebuilder:validation:Required
	Email string `json:"email" tf:"email"`

	// +kubebuilder:validation:Optional
	Phone *string `json:"phone,omitempty" tf:"phone"`
}

// SecurityCenterContactSpec defines the desired state of SecurityCenterContact
type SecurityCenterContactSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecurityCenterContactParameters `json:"forProvider"`
}

// SecurityCenterContactStatus defines the observed state of SecurityCenterContact.
type SecurityCenterContactStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecurityCenterContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterContact is the Schema for the SecurityCenterContacts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterContactSpec   `json:"spec"`
	Status            SecurityCenterContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterContactList contains a list of SecurityCenterContacts
type SecurityCenterContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterContact `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterContactKind             = "SecurityCenterContact"
	SecurityCenterContactGroupKind        = schema.GroupKind{Group: Group, Kind: SecurityCenterContactKind}.String()
	SecurityCenterContactKindAPIVersion   = SecurityCenterContactKind + "." + GroupVersion.String()
	SecurityCenterContactGroupVersionKind = GroupVersion.WithKind(SecurityCenterContactKind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterContact{}, &SecurityCenterContactList{})
}
