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

type HpcCacheBlobTargetObservation struct {
}

type HpcCacheBlobTargetParameters struct {

	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name"`

	// +kubebuilder:validation:Required
	CacheName string `json:"cacheName" tf:"cache_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NamespacePath string `json:"namespacePath" tf:"namespace_path"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	StorageContainerID string `json:"storageContainerId" tf:"storage_container_id"`
}

// HpcCacheBlobTargetSpec defines the desired state of HpcCacheBlobTarget
type HpcCacheBlobTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HpcCacheBlobTargetParameters `json:"forProvider"`
}

// HpcCacheBlobTargetStatus defines the observed state of HpcCacheBlobTarget.
type HpcCacheBlobTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HpcCacheBlobTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HpcCacheBlobTarget is the Schema for the HpcCacheBlobTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HpcCacheBlobTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HpcCacheBlobTargetSpec   `json:"spec"`
	Status            HpcCacheBlobTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HpcCacheBlobTargetList contains a list of HpcCacheBlobTargets
type HpcCacheBlobTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HpcCacheBlobTarget `json:"items"`
}

// Repository type metadata.
var (
	HpcCacheBlobTargetKind             = "HpcCacheBlobTarget"
	HpcCacheBlobTargetGroupKind        = schema.GroupKind{Group: Group, Kind: HpcCacheBlobTargetKind}.String()
	HpcCacheBlobTargetKindAPIVersion   = HpcCacheBlobTargetKind + "." + GroupVersion.String()
	HpcCacheBlobTargetGroupVersionKind = GroupVersion.WithKind(HpcCacheBlobTargetKind)
)

func init() {
	SchemeBuilder.Register(&HpcCacheBlobTarget{}, &HpcCacheBlobTargetList{})
}
