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

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Optional
	RequestHeader []RequestHeaderParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseHeader []ResponseHeaderParameters `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type FrontdoorRulesEngineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type FrontdoorRulesEngineParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	FrontdoorName *string `json:"frontdoorName" tf:"frontdoor_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Rule []FrontdoorRulesEngineRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FrontdoorRulesEngineRuleObservation struct {
}

type FrontdoorRulesEngineRuleParameters struct {

	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	MatchCondition []RuleMatchConditionParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`
}

type RequestHeaderObservation struct {
}

type RequestHeaderParameters struct {

	// +kubebuilder:validation:Optional
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResponseHeaderObservation struct {
}

type ResponseHeaderParameters struct {

	// +kubebuilder:validation:Optional
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuleMatchConditionObservation struct {
}

type RuleMatchConditionParameters struct {

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// +kubebuilder:validation:Optional
	Transform []*string `json:"transform,omitempty" tf:"transform,omitempty"`

	// +kubebuilder:validation:Optional
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`

	// +kubebuilder:validation:Optional
	Variable *string `json:"variable,omitempty" tf:"variable,omitempty"`
}

// FrontdoorRulesEngineSpec defines the desired state of FrontdoorRulesEngine
type FrontdoorRulesEngineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorRulesEngineParameters `json:"forProvider"`
}

// FrontdoorRulesEngineStatus defines the observed state of FrontdoorRulesEngine.
type FrontdoorRulesEngineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorRulesEngineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRulesEngine is the Schema for the FrontdoorRulesEngines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FrontdoorRulesEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorRulesEngineSpec   `json:"spec"`
	Status            FrontdoorRulesEngineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRulesEngineList contains a list of FrontdoorRulesEngines
type FrontdoorRulesEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorRulesEngine `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorRulesEngine_Kind             = "FrontdoorRulesEngine"
	FrontdoorRulesEngine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorRulesEngine_Kind}.String()
	FrontdoorRulesEngine_KindAPIVersion   = FrontdoorRulesEngine_Kind + "." + CRDGroupVersion.String()
	FrontdoorRulesEngine_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorRulesEngine_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorRulesEngine{}, &FrontdoorRulesEngineList{})
}
