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

type CaptureDescriptionObservation struct {
}

type CaptureDescriptionParameters struct {

	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination"`

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Required
	Encoding string `json:"encoding" tf:"encoding"`

	// +kubebuilder:validation:Optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds"`

	// +kubebuilder:validation:Optional
	SizeLimitInBytes *int64 `json:"sizeLimitInBytes,omitempty" tf:"size_limit_in_bytes"`

	// +kubebuilder:validation:Optional
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty" tf:"skip_empty_archives"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +kubebuilder:validation:Required
	ArchiveNameFormat string `json:"archiveNameFormat" tf:"archive_name_format"`

	// +kubebuilder:validation:Required
	BlobContainerName string `json:"blobContainerName" tf:"blob_container_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	StorageAccountID string `json:"storageAccountId" tf:"storage_account_id"`
}

type EventhubObservation struct {
	PartitionIds []string `json:"partitionIds,omitempty" tf:"partition_ids"`
}

type EventhubParameters struct {

	// +kubebuilder:validation:Optional
	CaptureDescription []CaptureDescriptionParameters `json:"captureDescription,omitempty" tf:"capture_description"`

	// +kubebuilder:validation:Required
	MessageRetention int64 `json:"messageRetention" tf:"message_retention"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NamespaceName string `json:"namespaceName" tf:"namespace_name"`

	// +kubebuilder:validation:Required
	PartitionCount int64 `json:"partitionCount" tf:"partition_count"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status"`
}

// EventhubSpec defines the desired state of Eventhub
type EventhubSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EventhubParameters `json:"forProvider"`
}

// EventhubStatus defines the observed state of Eventhub.
type EventhubStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EventhubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Eventhub is the Schema for the Eventhubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Eventhub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubSpec   `json:"spec"`
	Status            EventhubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubList contains a list of Eventhubs
type EventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eventhub `json:"items"`
}

// Repository type metadata.
var (
	EventhubKind             = "Eventhub"
	EventhubGroupKind        = schema.GroupKind{Group: Group, Kind: EventhubKind}.String()
	EventhubKindAPIVersion   = EventhubKind + "." + GroupVersion.String()
	EventhubGroupVersionKind = GroupVersion.WithKind(EventhubKind)
)

func init() {
	SchemeBuilder.Register(&Eventhub{}, &EventhubList{})
}
