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

type EndpointObservation struct {
}

type EndpointParameters struct {

	// +kubebuilder:validation:Optional
	BatchFrequencyInSeconds *int64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds"`

	// +kubebuilder:validation:Required
	ConnectionString string `json:"connectionString" tf:"connection_string"`

	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name"`

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`

	// +kubebuilder:validation:Optional
	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format"`

	// +kubebuilder:validation:Optional
	MaxChunkSizeInBytes *int64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type EnrichmentObservation struct {
}

type EnrichmentParameters struct {

	// +kubebuilder:validation:Required
	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`

	// +kubebuilder:validation:Required
	Key string `json:"key" tf:"key"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FallbackRouteObservation struct {
}

type FallbackRouteParameters struct {

	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Optional
	EndpointNames []string `json:"endpointNames,omitempty" tf:"endpoint_names"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source"`
}

type FileUploadObservation struct {
}

type FileUploadParameters struct {

	// +kubebuilder:validation:Required
	ConnectionString string `json:"connectionString" tf:"connection_string"`

	// +kubebuilder:validation:Required
	ContainerName string `json:"containerName" tf:"container_name"`

	// +kubebuilder:validation:Optional
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl"`

	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration"`

	// +kubebuilder:validation:Optional
	MaxDeliveryCount *int64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count"`

	// +kubebuilder:validation:Optional
	Notifications *bool `json:"notifications,omitempty" tf:"notifications"`

	// +kubebuilder:validation:Optional
	SasTTL *string `json:"sasTtl,omitempty" tf:"sas_ttl"`
}

type IPFilterRuleObservation struct {
}

type IPFilterRuleParameters struct {

	// +kubebuilder:validation:Required
	Action string `json:"action" tf:"action"`

	// +kubebuilder:validation:Required
	IPMask string `json:"ipMask" tf:"ip_mask"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type IothubObservation struct {
	EventHubEventsEndpoint string `json:"eventHubEventsEndpoint,omitempty" tf:"event_hub_events_endpoint"`

	EventHubEventsPath string `json:"eventHubEventsPath,omitempty" tf:"event_hub_events_path"`

	EventHubOperationsEndpoint string `json:"eventHubOperationsEndpoint,omitempty" tf:"event_hub_operations_endpoint"`

	EventHubOperationsPath string `json:"eventHubOperationsPath,omitempty" tf:"event_hub_operations_path"`

	Hostname string `json:"hostname,omitempty" tf:"hostname"`

	SharedAccessPolicy []SharedAccessPolicyObservation `json:"sharedAccessPolicy,omitempty" tf:"shared_access_policy"`

	Type string `json:"type,omitempty" tf:"type"`
}

type IothubParameters struct {

	// +kubebuilder:validation:Optional
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint"`

	// +kubebuilder:validation:Optional
	Enrichment []EnrichmentParameters `json:"enrichment,omitempty" tf:"enrichment"`

	// +kubebuilder:validation:Optional
	EventHubPartitionCount *int64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count"`

	// +kubebuilder:validation:Optional
	EventHubRetentionInDays *int64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days"`

	// +kubebuilder:validation:Optional
	FallbackRoute []FallbackRouteParameters `json:"fallbackRoute,omitempty" tf:"fallback_route"`

	// +kubebuilder:validation:Optional
	FileUpload []FileUploadParameters `json:"fileUpload,omitempty" tf:"file_upload"`

	// +kubebuilder:validation:Optional
	IPFilterRule []IPFilterRuleParameters `json:"ipFilterRule,omitempty" tf:"ip_filter_rule"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type RouteObservation struct {
}

type RouteParameters struct {

	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition"`

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Required
	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Source string `json:"source" tf:"source"`
}

type SharedAccessPolicyObservation struct {
	KeyName string `json:"keyName,omitempty" tf:"key_name"`

	Permissions string `json:"permissions,omitempty" tf:"permissions"`

	PrimaryKey string `json:"primaryKey,omitempty" tf:"primary_key"`

	SecondaryKey string `json:"secondaryKey,omitempty" tf:"secondary_key"`
}

type SharedAccessPolicyParameters struct {
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity int64 `json:"capacity" tf:"capacity"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

// IothubSpec defines the desired state of Iothub
type IothubSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IothubParameters `json:"forProvider"`
}

// IothubStatus defines the observed state of Iothub.
type IothubStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IothubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Iothub is the Schema for the Iothubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Iothub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSpec   `json:"spec"`
	Status            IothubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubList contains a list of Iothubs
type IothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Iothub `json:"items"`
}

// Repository type metadata.
var (
	IothubKind             = "Iothub"
	IothubGroupKind        = schema.GroupKind{Group: Group, Kind: IothubKind}.String()
	IothubKindAPIVersion   = IothubKind + "." + GroupVersion.String()
	IothubGroupVersionKind = GroupVersion.WithKind(IothubKind)
)

func init() {
	SchemeBuilder.Register(&Iothub{}, &IothubList{})
}
