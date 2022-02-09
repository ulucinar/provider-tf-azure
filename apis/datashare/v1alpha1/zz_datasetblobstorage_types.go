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

type DataSetBlobStorageObservation struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSetBlobStorageParameters struct {

	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Required
	DataShareID *string `json:"dataShareId" tf:"data_share_id,omitempty"`

	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// +kubebuilder:validation:Optional
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccount []StorageAccountParameters `json:"storageAccount" tf:"storage_account,omitempty"`
}

type StorageAccountObservation struct {
}

type StorageAccountParameters struct {

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
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`
}

// DataSetBlobStorageSpec defines the desired state of DataSetBlobStorage
type DataSetBlobStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetBlobStorageParameters `json:"forProvider"`
}

// DataSetBlobStorageStatus defines the observed state of DataSetBlobStorage.
type DataSetBlobStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetBlobStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetBlobStorage is the Schema for the DataSetBlobStorages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DataSetBlobStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetBlobStorageSpec   `json:"spec"`
	Status            DataSetBlobStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetBlobStorageList contains a list of DataSetBlobStorages
type DataSetBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetBlobStorage `json:"items"`
}

// Repository type metadata.
var (
	DataSetBlobStorage_Kind             = "DataSetBlobStorage"
	DataSetBlobStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetBlobStorage_Kind}.String()
	DataSetBlobStorage_KindAPIVersion   = DataSetBlobStorage_Kind + "." + CRDGroupVersion.String()
	DataSetBlobStorage_GroupVersionKind = CRDGroupVersion.WithKind(DataSetBlobStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetBlobStorage{}, &DataSetBlobStorageList{})
}
