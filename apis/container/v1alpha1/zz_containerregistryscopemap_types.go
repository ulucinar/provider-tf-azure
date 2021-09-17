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

type ContainerRegistryScopeMapObservation struct {
}

type ContainerRegistryScopeMapParameters struct {

	// +kubebuilder:validation:Required
	Actions []string `json:"actions" tf:"actions"`

	// +kubebuilder:validation:Required
	ContainerRegistryName string `json:"containerRegistryName" tf:"container_registry_name"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

// ContainerRegistryScopeMapSpec defines the desired state of ContainerRegistryScopeMap
type ContainerRegistryScopeMapSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ContainerRegistryScopeMapParameters `json:"forProvider"`
}

// ContainerRegistryScopeMapStatus defines the observed state of ContainerRegistryScopeMap.
type ContainerRegistryScopeMapStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ContainerRegistryScopeMapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRegistryScopeMap is the Schema for the ContainerRegistryScopeMaps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ContainerRegistryScopeMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerRegistryScopeMapSpec   `json:"spec"`
	Status            ContainerRegistryScopeMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRegistryScopeMapList contains a list of ContainerRegistryScopeMaps
type ContainerRegistryScopeMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerRegistryScopeMap `json:"items"`
}

// Repository type metadata.
var (
	ContainerRegistryScopeMapKind             = "ContainerRegistryScopeMap"
	ContainerRegistryScopeMapGroupKind        = schema.GroupKind{Group: Group, Kind: ContainerRegistryScopeMapKind}.String()
	ContainerRegistryScopeMapKindAPIVersion   = ContainerRegistryScopeMapKind + "." + GroupVersion.String()
	ContainerRegistryScopeMapGroupVersionKind = GroupVersion.WithKind(ContainerRegistryScopeMapKind)
)

func init() {
	SchemeBuilder.Register(&ContainerRegistryScopeMap{}, &ContainerRegistryScopeMapList{})
}
