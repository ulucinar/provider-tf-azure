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

type CodePackageObservation struct {
}

type CodePackageParameters struct {

	// +kubebuilder:validation:Required
	ImageName string `json:"imageName" tf:"image_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Resources []ResourcesParameters `json:"resources" tf:"resources"`
}

type LimitsObservation struct {
}

type LimitsParameters struct {

	// +kubebuilder:validation:Required
	CPU float64 `json:"cpu" tf:"cpu"`

	// +kubebuilder:validation:Required
	Memory float64 `json:"memory" tf:"memory"`
}

type RequestsObservation struct {
}

type RequestsParameters struct {

	// +kubebuilder:validation:Required
	CPU float64 `json:"cpu" tf:"cpu"`

	// +kubebuilder:validation:Required
	Memory float64 `json:"memory" tf:"memory"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	Limits []LimitsParameters `json:"limits,omitempty" tf:"limits"`

	// +kubebuilder:validation:Required
	Requests []RequestsParameters `json:"requests" tf:"requests"`
}

type ServiceFabricMeshApplicationObservation struct {
}

type ServiceFabricMeshApplicationParameters struct {

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Service []ServiceParameters `json:"service" tf:"service"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ServiceObservation struct {
}

type ServiceParameters struct {

	// +kubebuilder:validation:Required
	CodePackage []CodePackageParameters `json:"codePackage" tf:"code_package"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	OsType string `json:"osType" tf:"os_type"`
}

// ServiceFabricMeshApplicationSpec defines the desired state of ServiceFabricMeshApplication
type ServiceFabricMeshApplicationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceFabricMeshApplicationParameters `json:"forProvider"`
}

// ServiceFabricMeshApplicationStatus defines the observed state of ServiceFabricMeshApplication.
type ServiceFabricMeshApplicationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceFabricMeshApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceFabricMeshApplication is the Schema for the ServiceFabricMeshApplications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServiceFabricMeshApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceFabricMeshApplicationSpec   `json:"spec"`
	Status            ServiceFabricMeshApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceFabricMeshApplicationList contains a list of ServiceFabricMeshApplications
type ServiceFabricMeshApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceFabricMeshApplication `json:"items"`
}

// Repository type metadata.
var (
	ServiceFabricMeshApplicationKind             = "ServiceFabricMeshApplication"
	ServiceFabricMeshApplicationGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceFabricMeshApplicationKind}.String()
	ServiceFabricMeshApplicationKindAPIVersion   = ServiceFabricMeshApplicationKind + "." + GroupVersion.String()
	ServiceFabricMeshApplicationGroupVersionKind = GroupVersion.WithKind(ServiceFabricMeshApplicationKind)
)

func init() {
	SchemeBuilder.Register(&ServiceFabricMeshApplication{}, &ServiceFabricMeshApplicationList{})
}
