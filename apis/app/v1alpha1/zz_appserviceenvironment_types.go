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

type AppServiceEnvironmentObservation struct {
	InternalIPAddress string `json:"internalIpAddress,omitempty" tf:"internal_ip_address"`

	Location string `json:"location,omitempty" tf:"location"`

	OutboundIPAddresses []string `json:"outboundIpAddresses,omitempty" tf:"outbound_ip_addresses"`

	ServiceIPAddress string `json:"serviceIpAddress,omitempty" tf:"service_ip_address"`
}

type AppServiceEnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	AllowedUserIPCidrs []string `json:"allowedUserIpCidrs,omitempty" tf:"allowed_user_ip_cidrs"`

	// +kubebuilder:validation:Optional
	ClusterSetting []ClusterSettingParameters `json:"clusterSetting,omitempty" tf:"cluster_setting"`

	// +kubebuilder:validation:Optional
	FrontEndScaleFactor *int64 `json:"frontEndScaleFactor,omitempty" tf:"front_end_scale_factor"`

	// +kubebuilder:validation:Optional
	InternalLoadBalancingMode *string `json:"internalLoadBalancingMode,omitempty" tf:"internal_load_balancing_mode"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier"`

	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	UserWhitelistedIPRanges []string `json:"userWhitelistedIpRanges,omitempty" tf:"user_whitelisted_ip_ranges"`
}

type ClusterSettingObservation struct {
}

type ClusterSettingParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

// AppServiceEnvironmentSpec defines the desired state of AppServiceEnvironment
type AppServiceEnvironmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppServiceEnvironmentParameters `json:"forProvider"`
}

// AppServiceEnvironmentStatus defines the observed state of AppServiceEnvironment.
type AppServiceEnvironmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppServiceEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceEnvironment is the Schema for the AppServiceEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceEnvironmentSpec   `json:"spec"`
	Status            AppServiceEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceEnvironmentList contains a list of AppServiceEnvironments
type AppServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServiceEnvironment `json:"items"`
}

// Repository type metadata.
var (
	AppServiceEnvironmentKind             = "AppServiceEnvironment"
	AppServiceEnvironmentGroupKind        = schema.GroupKind{Group: Group, Kind: AppServiceEnvironmentKind}.String()
	AppServiceEnvironmentKindAPIVersion   = AppServiceEnvironmentKind + "." + GroupVersion.String()
	AppServiceEnvironmentGroupVersionKind = GroupVersion.WithKind(AppServiceEnvironmentKind)
)

func init() {
	SchemeBuilder.Register(&AppServiceEnvironment{}, &AppServiceEnvironmentList{})
}
