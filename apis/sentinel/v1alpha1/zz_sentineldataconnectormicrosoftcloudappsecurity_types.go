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

type SentinelDataConnectorMicrosoftCloudAppSecurityObservation struct {
}

type SentinelDataConnectorMicrosoftCloudAppSecurityParameters struct {

	// +kubebuilder:validation:Optional
	AlertsEnabled *bool `json:"alertsEnabled,omitempty" tf:"alerts_enabled"`

	// +kubebuilder:validation:Optional
	DiscoveryLogsEnabled *bool `json:"discoveryLogsEnabled,omitempty" tf:"discovery_logs_enabled"`

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

// SentinelDataConnectorMicrosoftCloudAppSecuritySpec defines the desired state of SentinelDataConnectorMicrosoftCloudAppSecurity
type SentinelDataConnectorMicrosoftCloudAppSecuritySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SentinelDataConnectorMicrosoftCloudAppSecurityParameters `json:"forProvider"`
}

// SentinelDataConnectorMicrosoftCloudAppSecurityStatus defines the observed state of SentinelDataConnectorMicrosoftCloudAppSecurity.
type SentinelDataConnectorMicrosoftCloudAppSecurityStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SentinelDataConnectorMicrosoftCloudAppSecurityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorMicrosoftCloudAppSecurity is the Schema for the SentinelDataConnectorMicrosoftCloudAppSecuritys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelDataConnectorMicrosoftCloudAppSecurity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelDataConnectorMicrosoftCloudAppSecuritySpec   `json:"spec"`
	Status            SentinelDataConnectorMicrosoftCloudAppSecurityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelDataConnectorMicrosoftCloudAppSecurityList contains a list of SentinelDataConnectorMicrosoftCloudAppSecuritys
type SentinelDataConnectorMicrosoftCloudAppSecurityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelDataConnectorMicrosoftCloudAppSecurity `json:"items"`
}

// Repository type metadata.
var (
	SentinelDataConnectorMicrosoftCloudAppSecurityKind             = "SentinelDataConnectorMicrosoftCloudAppSecurity"
	SentinelDataConnectorMicrosoftCloudAppSecurityGroupKind        = schema.GroupKind{Group: Group, Kind: SentinelDataConnectorMicrosoftCloudAppSecurityKind}.String()
	SentinelDataConnectorMicrosoftCloudAppSecurityKindAPIVersion   = SentinelDataConnectorMicrosoftCloudAppSecurityKind + "." + GroupVersion.String()
	SentinelDataConnectorMicrosoftCloudAppSecurityGroupVersionKind = GroupVersion.WithKind(SentinelDataConnectorMicrosoftCloudAppSecurityKind)
)

func init() {
	SchemeBuilder.Register(&SentinelDataConnectorMicrosoftCloudAppSecurity{}, &SentinelDataConnectorMicrosoftCloudAppSecurityList{})
}
