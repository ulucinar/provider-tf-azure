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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ModuleObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ModuleParameters struct {

	// +kubebuilder:validation:Optional
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type RedisEnterpriseDatabaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RedisEnterpriseDatabaseParameters struct {

	// +kubebuilder:validation:Optional
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol,omitempty"`

	// +crossplane:generate:reference:type=RedisEnterpriseCluster
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusteringPolicy *string `json:"clusteringPolicy,omitempty" tf:"clustering_policy,omitempty"`

	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Module []ModuleParameters `json:"module,omitempty" tf:"module,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// RedisEnterpriseDatabaseSpec defines the desired state of RedisEnterpriseDatabase
type RedisEnterpriseDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisEnterpriseDatabaseParameters `json:"forProvider"`
}

// RedisEnterpriseDatabaseStatus defines the observed state of RedisEnterpriseDatabase.
type RedisEnterpriseDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisEnterpriseDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseDatabase is the Schema for the RedisEnterpriseDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RedisEnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseDatabaseSpec   `json:"spec"`
	Status            RedisEnterpriseDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseDatabaseList contains a list of RedisEnterpriseDatabases
type RedisEnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseDatabase `json:"items"`
}

// Repository type metadata.
var (
	RedisEnterpriseDatabase_Kind             = "RedisEnterpriseDatabase"
	RedisEnterpriseDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisEnterpriseDatabase_Kind}.String()
	RedisEnterpriseDatabase_KindAPIVersion   = RedisEnterpriseDatabase_Kind + "." + CRDGroupVersion.String()
	RedisEnterpriseDatabase_GroupVersionKind = CRDGroupVersion.WithKind(RedisEnterpriseDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisEnterpriseDatabase{}, &RedisEnterpriseDatabaseList{})
}
