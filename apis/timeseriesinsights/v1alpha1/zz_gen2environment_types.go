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

type Gen2EnvironmentObservation struct {
	DataAccessFqdn *string `json:"dataAccessFqdn,omitempty" tf:"data_access_fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type Gen2EnvironmentParameters struct {

	// +kubebuilder:validation:Required
	IDProperties []*string `json:"idProperties" tf:"id_properties,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Required
	Storage []StorageParameters `json:"storage" tf:"storage,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WarmStoreDataRetentionTime *string `json:"warmStoreDataRetentionTime,omitempty" tf:"warm_store_data_retention_time,omitempty"`
}

type StorageObservation struct {
}

type StorageParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// Gen2EnvironmentSpec defines the desired state of Gen2Environment
type Gen2EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Gen2EnvironmentParameters `json:"forProvider"`
}

// Gen2EnvironmentStatus defines the observed state of Gen2Environment.
type Gen2EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Gen2EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gen2Environment is the Schema for the Gen2Environments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Gen2Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Gen2EnvironmentSpec   `json:"spec"`
	Status            Gen2EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Gen2EnvironmentList contains a list of Gen2Environments
type Gen2EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gen2Environment `json:"items"`
}

// Repository type metadata.
var (
	Gen2Environment_Kind             = "Gen2Environment"
	Gen2Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gen2Environment_Kind}.String()
	Gen2Environment_KindAPIVersion   = Gen2Environment_Kind + "." + CRDGroupVersion.String()
	Gen2Environment_GroupVersionKind = CRDGroupVersion.WithKind(Gen2Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Gen2Environment{}, &Gen2EnvironmentList{})
}
