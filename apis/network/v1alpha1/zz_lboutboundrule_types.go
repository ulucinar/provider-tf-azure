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

type LBOutboundRuleFrontendIPConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBOutboundRuleFrontendIPConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LBOutboundRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBOutboundRuleParameters struct {

	// +kubebuilder:validation:Optional
	AllocatedOutboundPorts *int64 `json:"allocatedOutboundPorts,omitempty" tf:"allocated_outbound_ports,omitempty"`

	// +kubebuilder:validation:Required
	BackendAddressPoolID *string `json:"backendAddressPoolId" tf:"backend_address_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// +kubebuilder:validation:Optional
	FrontendIPConfiguration []LBOutboundRuleFrontendIPConfigurationParameters `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	LoadbalancerID *string `json:"loadbalancerId" tf:"loadbalancer_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// LBOutboundRuleSpec defines the desired state of LBOutboundRule
type LBOutboundRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBOutboundRuleParameters `json:"forProvider"`
}

// LBOutboundRuleStatus defines the observed state of LBOutboundRule.
type LBOutboundRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBOutboundRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBOutboundRule is the Schema for the LBOutboundRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LBOutboundRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBOutboundRuleSpec   `json:"spec"`
	Status            LBOutboundRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBOutboundRuleList contains a list of LBOutboundRules
type LBOutboundRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBOutboundRule `json:"items"`
}

// Repository type metadata.
var (
	LBOutboundRule_Kind             = "LBOutboundRule"
	LBOutboundRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBOutboundRule_Kind}.String()
	LBOutboundRule_KindAPIVersion   = LBOutboundRule_Kind + "." + CRDGroupVersion.String()
	LBOutboundRule_GroupVersionKind = CRDGroupVersion.WithKind(LBOutboundRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LBOutboundRule{}, &LBOutboundRuleList{})
}
