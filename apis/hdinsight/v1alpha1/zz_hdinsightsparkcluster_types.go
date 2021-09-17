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

type HdinsightSparkClusterComponentVersionObservation struct {
}

type HdinsightSparkClusterComponentVersionParameters struct {

	// +kubebuilder:validation:Required
	Spark string `json:"spark" tf:"spark"`
}

type HdinsightSparkClusterGatewayObservation struct {
}

type HdinsightSparkClusterGatewayParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresAmbariObservation struct {
}

type HdinsightSparkClusterMetastoresAmbariParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName" tf:"database_name"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Server string `json:"server" tf:"server"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresHiveObservation struct {
}

type HdinsightSparkClusterMetastoresHiveParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName" tf:"database_name"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Server string `json:"server" tf:"server"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresObservation struct {
}

type HdinsightSparkClusterMetastoresOozieObservation struct {
}

type HdinsightSparkClusterMetastoresOozieParameters struct {

	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName" tf:"database_name"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Server string `json:"server" tf:"server"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterMetastoresParameters struct {

	// +kubebuilder:validation:Optional
	Ambari []HdinsightSparkClusterMetastoresAmbariParameters `json:"ambari,omitempty" tf:"ambari"`

	// +kubebuilder:validation:Optional
	Hive []HdinsightSparkClusterMetastoresHiveParameters `json:"hive,omitempty" tf:"hive"`

	// +kubebuilder:validation:Optional
	Oozie []HdinsightSparkClusterMetastoresOozieParameters `json:"oozie,omitempty" tf:"oozie"`
}

type HdinsightSparkClusterMonitorObservation struct {
}

type HdinsightSparkClusterMonitorParameters struct {

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	// +kubebuilder:validation:Required
	PrimaryKey string `json:"primaryKey" tf:"primary_key"`
}

type HdinsightSparkClusterNetworkObservation struct {
}

type HdinsightSparkClusterNetworkParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction"`

	// +kubebuilder:validation:Optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`
}

type HdinsightSparkClusterObservation struct {
	HTTPSEndpoint string `json:"httpsEndpoint,omitempty" tf:"https_endpoint"`

	SSHEndpoint string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint"`
}

type HdinsightSparkClusterParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	// +kubebuilder:validation:Required
	ComponentVersion []HdinsightSparkClusterComponentVersionParameters `json:"componentVersion" tf:"component_version"`

	// +kubebuilder:validation:Optional
	EncryptionInTransitEnabled *bool `json:"encryptionInTransitEnabled,omitempty" tf:"encryption_in_transit_enabled"`

	// +kubebuilder:validation:Required
	Gateway []HdinsightSparkClusterGatewayParameters `json:"gateway" tf:"gateway"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Optional
	Metastores []HdinsightSparkClusterMetastoresParameters `json:"metastores,omitempty" tf:"metastores"`

	// +kubebuilder:validation:Optional
	Monitor []HdinsightSparkClusterMonitorParameters `json:"monitor,omitempty" tf:"monitor"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Network []HdinsightSparkClusterNetworkParameters `json:"network,omitempty" tf:"network"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Roles []HdinsightSparkClusterRolesParameters `json:"roles" tf:"roles"`

	// +kubebuilder:validation:Optional
	StorageAccount []HdinsightSparkClusterStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account"`

	// +kubebuilder:validation:Optional
	StorageAccountGen2 []HdinsightSparkClusterStorageAccountGen2Parameters `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2"`

	// +kubebuilder:validation:Optional
	TLSMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	Tier string `json:"tier" tf:"tier"`
}

type HdinsightSparkClusterRolesHeadNodeObservation struct {
}

type HdinsightSparkClusterRolesHeadNodeParameters struct {

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

type HdinsightSparkClusterRolesObservation struct {
}

type HdinsightSparkClusterRolesParameters struct {

	// +kubebuilder:validation:Required
	HeadNode []HdinsightSparkClusterRolesHeadNodeParameters `json:"headNode" tf:"head_node"`

	// +kubebuilder:validation:Required
	WorkerNode []HdinsightSparkClusterRolesWorkerNodeParameters `json:"workerNode" tf:"worker_node"`

	// +kubebuilder:validation:Required
	ZookeeperNode []HdinsightSparkClusterRolesZookeeperNodeParameters `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightSparkClusterRolesWorkerNodeAutoscaleObservation struct {
}

type HdinsightSparkClusterRolesWorkerNodeAutoscaleParameters struct {

	// +kubebuilder:validation:Optional
	Capacity []WorkerNodeAutoscaleCapacityParameters `json:"capacity,omitempty" tf:"capacity"`

	// +kubebuilder:validation:Optional
	Recurrence []RolesWorkerNodeAutoscaleRecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence"`
}

type HdinsightSparkClusterRolesWorkerNodeObservation struct {
}

type HdinsightSparkClusterRolesWorkerNodeParameters struct {

	// +kubebuilder:validation:Optional
	Autoscale []HdinsightSparkClusterRolesWorkerNodeAutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale"`

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

type HdinsightSparkClusterRolesZookeeperNodeObservation struct {
}

type HdinsightSparkClusterRolesZookeeperNodeParameters struct {

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

type HdinsightSparkClusterStorageAccountGen2Observation struct {
}

type HdinsightSparkClusterStorageAccountGen2Parameters struct {

	// +kubebuilder:validation:Required
	FilesystemID string `json:"filesystemId" tf:"filesystem_id"`

	// +kubebuilder:validation:Required
	IsDefault bool `json:"isDefault" tf:"is_default"`

	// +kubebuilder:validation:Required
	ManagedIdentityResourceID string `json:"managedIdentityResourceId" tf:"managed_identity_resource_id"`

	// +kubebuilder:validation:Required
	StorageResourceID string `json:"storageResourceId" tf:"storage_resource_id"`
}

type HdinsightSparkClusterStorageAccountObservation struct {
}

type HdinsightSparkClusterStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	IsDefault bool `json:"isDefault" tf:"is_default"`

	// +kubebuilder:validation:Required
	StorageAccountKey string `json:"storageAccountKey" tf:"storage_account_key"`

	// +kubebuilder:validation:Required
	StorageContainerID string `json:"storageContainerId" tf:"storage_container_id"`
}

type RolesWorkerNodeAutoscaleRecurrenceObservation struct {
}

type RolesWorkerNodeAutoscaleRecurrenceParameters struct {

	// +kubebuilder:validation:Required
	Schedule []WorkerNodeAutoscaleRecurrenceScheduleParameters `json:"schedule" tf:"schedule"`

	// +kubebuilder:validation:Required
	Timezone string `json:"timezone" tf:"timezone"`
}

type WorkerNodeAutoscaleCapacityObservation struct {
}

type WorkerNodeAutoscaleCapacityParameters struct {

	// +kubebuilder:validation:Required
	MaxInstanceCount int64 `json:"maxInstanceCount" tf:"max_instance_count"`

	// +kubebuilder:validation:Required
	MinInstanceCount int64 `json:"minInstanceCount" tf:"min_instance_count"`
}

type WorkerNodeAutoscaleRecurrenceScheduleObservation struct {
}

type WorkerNodeAutoscaleRecurrenceScheduleParameters struct {

	// +kubebuilder:validation:Required
	Days []string `json:"days" tf:"days"`

	// +kubebuilder:validation:Required
	TargetInstanceCount int64 `json:"targetInstanceCount" tf:"target_instance_count"`

	// +kubebuilder:validation:Required
	Time string `json:"time" tf:"time"`
}

// HdinsightSparkClusterSpec defines the desired state of HdinsightSparkCluster
type HdinsightSparkClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HdinsightSparkClusterParameters `json:"forProvider"`
}

// HdinsightSparkClusterStatus defines the observed state of HdinsightSparkCluster.
type HdinsightSparkClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HdinsightSparkClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightSparkCluster is the Schema for the HdinsightSparkClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HdinsightSparkCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightSparkClusterSpec   `json:"spec"`
	Status            HdinsightSparkClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HdinsightSparkClusterList contains a list of HdinsightSparkClusters
type HdinsightSparkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HdinsightSparkCluster `json:"items"`
}

// Repository type metadata.
var (
	HdinsightSparkClusterKind             = "HdinsightSparkCluster"
	HdinsightSparkClusterGroupKind        = schema.GroupKind{Group: Group, Kind: HdinsightSparkClusterKind}.String()
	HdinsightSparkClusterKindAPIVersion   = HdinsightSparkClusterKind + "." + GroupVersion.String()
	HdinsightSparkClusterGroupVersionKind = GroupVersion.WithKind(HdinsightSparkClusterKind)
)

func init() {
	SchemeBuilder.Register(&HdinsightSparkCluster{}, &HdinsightSparkClusterList{})
}
