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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VariableIntObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VariableIntParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Value *int64 `json:"value,omitempty" tf:"value,omitempty"`
}

// VariableIntSpec defines the desired state of VariableInt
type VariableIntSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VariableIntParameters `json:"forProvider"`
}

// VariableIntStatus defines the observed state of VariableInt.
type VariableIntStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VariableIntObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VariableInt is the Schema for the VariableInts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VariableInt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VariableIntSpec   `json:"spec"`
	Status            VariableIntStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableIntList contains a list of VariableInts
type VariableIntList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VariableInt `json:"items"`
}

// Repository type metadata.
var (
	VariableInt_Kind             = "VariableInt"
	VariableInt_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VariableInt_Kind}.String()
	VariableInt_KindAPIVersion   = VariableInt_Kind + "." + CRDGroupVersion.String()
	VariableInt_GroupVersionKind = CRDGroupVersion.WithKind(VariableInt_Kind)
)

func init() {
	SchemeBuilder.Register(&VariableInt{}, &VariableIntList{})
}
