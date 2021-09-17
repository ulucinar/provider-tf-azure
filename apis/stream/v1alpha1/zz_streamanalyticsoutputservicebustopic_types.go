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

type StreamAnalyticsOutputServicebusTopicObservation struct {
}

type StreamAnalyticsOutputServicebusTopicParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Serialization []StreamAnalyticsOutputServicebusTopicSerializationParameters `json:"serialization" tf:"serialization"`

	// +kubebuilder:validation:Required
	ServicebusNamespace string `json:"servicebusNamespace" tf:"servicebus_namespace"`

	// +kubebuilder:validation:Required
	SharedAccessPolicyKey string `json:"sharedAccessPolicyKey" tf:"shared_access_policy_key"`

	// +kubebuilder:validation:Required
	SharedAccessPolicyName string `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`

	// +kubebuilder:validation:Required
	StreamAnalyticsJobName string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`

	// +kubebuilder:validation:Required
	TopicName string `json:"topicName" tf:"topic_name"`
}

type StreamAnalyticsOutputServicebusTopicSerializationObservation struct {
}

type StreamAnalyticsOutputServicebusTopicSerializationParameters struct {

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`

	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

// StreamAnalyticsOutputServicebusTopicSpec defines the desired state of StreamAnalyticsOutputServicebusTopic
type StreamAnalyticsOutputServicebusTopicSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StreamAnalyticsOutputServicebusTopicParameters `json:"forProvider"`
}

// StreamAnalyticsOutputServicebusTopicStatus defines the observed state of StreamAnalyticsOutputServicebusTopic.
type StreamAnalyticsOutputServicebusTopicStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StreamAnalyticsOutputServicebusTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamAnalyticsOutputServicebusTopic is the Schema for the StreamAnalyticsOutputServicebusTopics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StreamAnalyticsOutputServicebusTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputServicebusTopicSpec   `json:"spec"`
	Status            StreamAnalyticsOutputServicebusTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamAnalyticsOutputServicebusTopicList contains a list of StreamAnalyticsOutputServicebusTopics
type StreamAnalyticsOutputServicebusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamAnalyticsOutputServicebusTopic `json:"items"`
}

// Repository type metadata.
var (
	StreamAnalyticsOutputServicebusTopicKind             = "StreamAnalyticsOutputServicebusTopic"
	StreamAnalyticsOutputServicebusTopicGroupKind        = schema.GroupKind{Group: Group, Kind: StreamAnalyticsOutputServicebusTopicKind}.String()
	StreamAnalyticsOutputServicebusTopicKindAPIVersion   = StreamAnalyticsOutputServicebusTopicKind + "." + GroupVersion.String()
	StreamAnalyticsOutputServicebusTopicGroupVersionKind = GroupVersion.WithKind(StreamAnalyticsOutputServicebusTopicKind)
)

func init() {
	SchemeBuilder.Register(&StreamAnalyticsOutputServicebusTopic{}, &StreamAnalyticsOutputServicebusTopicList{})
}
