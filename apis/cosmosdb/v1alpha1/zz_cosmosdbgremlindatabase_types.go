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

type CosmosdbGremlinDatabaseAutoscaleSettingsObservation struct {
}

type CosmosdbGremlinDatabaseAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput"`
}

type CosmosdbGremlinDatabaseObservation struct {
}

type CosmosdbGremlinDatabaseParameters struct {

	// +kubebuilder:validation:Required
	AccountName string `json:"accountName" tf:"account_name"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []CosmosdbGremlinDatabaseAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
}

// CosmosdbGremlinDatabaseSpec defines the desired state of CosmosdbGremlinDatabase
type CosmosdbGremlinDatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CosmosdbGremlinDatabaseParameters `json:"forProvider"`
}

// CosmosdbGremlinDatabaseStatus defines the observed state of CosmosdbGremlinDatabase.
type CosmosdbGremlinDatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CosmosdbGremlinDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbGremlinDatabase is the Schema for the CosmosdbGremlinDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CosmosdbGremlinDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbGremlinDatabaseSpec   `json:"spec"`
	Status            CosmosdbGremlinDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbGremlinDatabaseList contains a list of CosmosdbGremlinDatabases
type CosmosdbGremlinDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosdbGremlinDatabase `json:"items"`
}

// Repository type metadata.
var (
	CosmosdbGremlinDatabaseKind             = "CosmosdbGremlinDatabase"
	CosmosdbGremlinDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: CosmosdbGremlinDatabaseKind}.String()
	CosmosdbGremlinDatabaseKindAPIVersion   = CosmosdbGremlinDatabaseKind + "." + GroupVersion.String()
	CosmosdbGremlinDatabaseGroupVersionKind = GroupVersion.WithKind(CosmosdbGremlinDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&CosmosdbGremlinDatabase{}, &CosmosdbGremlinDatabaseList{})
}
