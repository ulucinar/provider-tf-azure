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

type DataFactoryDatasetSnowflakeObservation struct {
}

type DataFactoryDatasetSnowflakeParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`

	// +kubebuilder:validation:Optional
	Annotations []string `json:"annotations,omitempty" tf:"annotations"`

	// +kubebuilder:validation:Required
	DataFactoryName string `json:"dataFactoryName" tf:"data_factory_name"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder"`

	// +kubebuilder:validation:Required
	LinkedServiceName string `json:"linkedServiceName" tf:"linked_service_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SchemaColumn []DataFactoryDatasetSnowflakeSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column"`

	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name"`

	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name"`
}

type DataFactoryDatasetSnowflakeSchemaColumnObservation struct {
}

type DataFactoryDatasetSnowflakeSchemaColumnParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

// DataFactoryDatasetSnowflakeSpec defines the desired state of DataFactoryDatasetSnowflake
type DataFactoryDatasetSnowflakeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataFactoryDatasetSnowflakeParameters `json:"forProvider"`
}

// DataFactoryDatasetSnowflakeStatus defines the observed state of DataFactoryDatasetSnowflake.
type DataFactoryDatasetSnowflakeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataFactoryDatasetSnowflakeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetSnowflake is the Schema for the DataFactoryDatasetSnowflakes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryDatasetSnowflake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryDatasetSnowflakeSpec   `json:"spec"`
	Status            DataFactoryDatasetSnowflakeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetSnowflakeList contains a list of DataFactoryDatasetSnowflakes
type DataFactoryDatasetSnowflakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryDatasetSnowflake `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryDatasetSnowflakeKind             = "DataFactoryDatasetSnowflake"
	DataFactoryDatasetSnowflakeGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryDatasetSnowflakeKind}.String()
	DataFactoryDatasetSnowflakeKindAPIVersion   = DataFactoryDatasetSnowflakeKind + "." + GroupVersion.String()
	DataFactoryDatasetSnowflakeGroupVersionKind = GroupVersion.WithKind(DataFactoryDatasetSnowflakeKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryDatasetSnowflake{}, &DataFactoryDatasetSnowflakeList{})
}
