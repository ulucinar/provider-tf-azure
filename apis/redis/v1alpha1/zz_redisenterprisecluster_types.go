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

type RedisEnterpriseClusterObservation struct {
	Hostname string `json:"hostname,omitempty" tf:"hostname"`

	Version string `json:"version,omitempty" tf:"version"`
}

type RedisEnterpriseClusterParameters struct {

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	SkuName string `json:"skuName" tf:"sku_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

// RedisEnterpriseClusterSpec defines the desired state of RedisEnterpriseCluster
type RedisEnterpriseClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RedisEnterpriseClusterParameters `json:"forProvider"`
}

// RedisEnterpriseClusterStatus defines the observed state of RedisEnterpriseCluster.
type RedisEnterpriseClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RedisEnterpriseClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseCluster is the Schema for the RedisEnterpriseClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisEnterpriseCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseClusterSpec   `json:"spec"`
	Status            RedisEnterpriseClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseClusterList contains a list of RedisEnterpriseClusters
type RedisEnterpriseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseCluster `json:"items"`
}

// Repository type metadata.
var (
	RedisEnterpriseClusterKind             = "RedisEnterpriseCluster"
	RedisEnterpriseClusterGroupKind        = schema.GroupKind{Group: Group, Kind: RedisEnterpriseClusterKind}.String()
	RedisEnterpriseClusterKindAPIVersion   = RedisEnterpriseClusterKind + "." + GroupVersion.String()
	RedisEnterpriseClusterGroupVersionKind = GroupVersion.WithKind(RedisEnterpriseClusterKind)
)

func init() {
	SchemeBuilder.Register(&RedisEnterpriseCluster{}, &RedisEnterpriseClusterList{})
}
