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

type HdinsightRserverClusterGatewayObservation struct {
}

type HdinsightRserverClusterGatewayParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

type HdinsightRserverClusterObservation struct {
	EdgeSSHEndpoint string `json:"edgeSshEndpoint,omitempty" tf:"edge_ssh_endpoint"`

	HTTPSEndpoint string `json:"httpsEndpoint,omitempty" tf:"https_endpoint"`

	SSHEndpoint string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint"`
}

type HdinsightRserverClusterParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	// +kubebuilder:validation:Required
	Gateway []HdinsightRserverClusterGatewayParameters `json:"gateway" tf:"gateway"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Roles []HdinsightRserverClusterRolesParameters `json:"roles" tf:"roles"`

	// +kubebuilder:validation:Required
	Rstudio bool `json:"rstudio" tf:"rstudio"`

	// +kubebuilder:validation:Optional
	StorageAccount []HdinsightRserverClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account"`

	// +kubebuilder:validation:Optional
	TLSMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	Tier string `json:"tier" tf:"tier"`
}

type HdinsightRserverClusterRolesEdgeNodeObservation struct {
}

type HdinsightRserverClusterRolesEdgeNodeParameters struct {

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password"`

	// +kubebuilder:validation:Optional
	SSHKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`

	// +kubebuilder:validation:Required
	VMSize string `json:"vmSize" tf:"vm_size"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`
}

type HdinsightRserverClusterRolesHeadNodeObservation struct {
}

type HdinsightRserverClusterRolesHeadNodeParameters struct {

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password"`

	// +kubebuilder:validation:Optional
	SSHKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`

	// +kubebuilder:validation:Required
	VMSize string `json:"vmSize" tf:"vm_size"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`
}

type HdinsightRserverClusterRolesObservation struct {
}

type HdinsightRserverClusterRolesParameters struct {

	// +kubebuilder:validation:Required
	EdgeNode []HdinsightRserverClusterRolesEdgeNodeParameters `json:"edgeNode" tf:"edge_node"`

	// +kubebuilder:validation:Required
	HeadNode []HdinsightRserverClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node"`

	// +kubebuilder:validation:Required
	WorkerNode []HdinsightRserverClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node"`

	// +kubebuilder:validation:Required
	ZookeeperNode []HdinsightRserverClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightRserverClusterRolesWorkerNodeObservation struct {
}

type HdinsightRserverClusterRolesWorkerNodeParameters struct {

	// +kubebuilder:validation:Optional
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count"`

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password"`

	// +kubebuilder:validation:Optional
	SSHKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// +kubebuilder:validation:Required
	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`

	// +kubebuilder:validation:Required
	VMSize string `json:"vmSize" tf:"vm_size"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`
}

type HdinsightRserverClusterRolesZookeeperNodeObservation struct {
}

type HdinsightRserverClusterRolesZookeeperNodeParameters struct {

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password"`

	// +kubebuilder:validation:Optional
	SSHKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`

	// +kubebuilder:validation:Required
	VMSize string `json:"vmSize" tf:"vm_size"`

	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id"`
}

type HdinsightRserverClusterStorageAccountObservation struct {
}

type HdinsightRserverClusterStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	IsDefault bool `json:"isDefault" tf:"is_default"`

	// +kubebuilder:validation:Required
	StorageAccountKey string `json:"storageAccountKey" tf:"storage_account_key"`

	// +kubebuilder:validation:Required
	StorageContainerID string `json:"storageContainerId" tf:"storage_container_id"`
}

// HdinsightRserverClusterSpec defines the desired state of HdinsightRserverCluster
type HdinsightRserverClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HdinsightRserverClusterParameters `json:"forProvider"`
}

// HdinsightRserverClusterStatus defines the observed state of HdinsightRserverCluster.
type HdinsightRserverClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HdinsightRserverClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightRserverCluster is the Schema for the HdinsightRserverClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HdinsightRserverCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightRserverClusterSpec   `json:"spec"`
	Status            HdinsightRserverClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightRserverClusterList contains a list of HdinsightRserverClusters
type HdinsightRserverClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HdinsightRserverCluster `json:"items"`
}

// Repository type metadata.
var (
	HdinsightRserverClusterKind             = "HdinsightRserverCluster"
	HdinsightRserverClusterGroupKind        = schema.GroupKind{Group: Group, Kind: HdinsightRserverClusterKind}.String()
	HdinsightRserverClusterKindAPIVersion   = HdinsightRserverClusterKind + "." + GroupVersion.String()
	HdinsightRserverClusterGroupVersionKind = GroupVersion.WithKind(HdinsightRserverClusterKind)
)

func init() {
	SchemeBuilder.Register(&HdinsightRserverCluster{}, &HdinsightRserverClusterList{})
}
