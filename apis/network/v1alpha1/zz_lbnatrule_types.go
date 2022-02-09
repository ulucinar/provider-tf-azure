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

type LBNATRuleObservation struct {
	BackendIPConfigurationID *string `json:"backendIpConfigurationId,omitempty" tf:"backend_ip_configuration_id,omitempty"`

	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBNATRuleParameters struct {

	// +kubebuilder:validation:Required
	BackendPort *int64 `json:"backendPort" tf:"backend_port,omitempty"`

	// +kubebuilder:validation:Optional
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// +kubebuilder:validation:Required
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName" tf:"frontend_ip_configuration_name,omitempty"`

	// +kubebuilder:validation:Required
	FrontendPort *int64 `json:"frontendPort" tf:"frontend_port,omitempty"`

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

// LBNATRuleSpec defines the desired state of LBNATRule
type LBNATRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBNATRuleParameters `json:"forProvider"`
}

// LBNATRuleStatus defines the observed state of LBNATRule.
type LBNATRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBNATRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBNATRule is the Schema for the LBNATRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LBNATRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBNATRuleSpec   `json:"spec"`
	Status            LBNATRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBNATRuleList contains a list of LBNATRules
type LBNATRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBNATRule `json:"items"`
}

// Repository type metadata.
var (
	LBNATRule_Kind             = "LBNATRule"
	LBNATRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBNATRule_Kind}.String()
	LBNATRule_KindAPIVersion   = LBNATRule_Kind + "." + CRDGroupVersion.String()
	LBNATRule_GroupVersionKind = CRDGroupVersion.WithKind(LBNATRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LBNATRule{}, &LBNATRuleList{})
}
