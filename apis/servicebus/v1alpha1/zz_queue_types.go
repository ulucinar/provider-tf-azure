/*
Copyright 2022 The Crossplane Authors.

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

type QueueObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QueueParameters struct {

	// +kubebuilder:validation:Optional
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// +kubebuilder:validation:Optional
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// +kubebuilder:validation:Optional
	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// +kubebuilder:validation:Optional
	MaxMessageSizeInKilobytes *float64 `json:"maxMessageSizeInKilobytes,omitempty" tf:"max_message_size_in_kilobytes,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSizeInMegabytes *float64 `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Optional
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`

	// +kubebuilder:validation:Optional
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
