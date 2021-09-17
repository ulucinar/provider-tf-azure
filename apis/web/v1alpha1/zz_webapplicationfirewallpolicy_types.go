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

type CustomRulesObservation struct {
}

type CustomRulesParameters struct {

	// +kubebuilder:validation:Required
	Action string `json:"action" tf:"action"`

	// +kubebuilder:validation:Required
	MatchConditions []MatchConditionsParameters `json:"matchConditions" tf:"match_conditions"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Required
	Priority int64 `json:"priority" tf:"priority"`

	// +kubebuilder:validation:Required
	RuleType string `json:"ruleType" tf:"rule_type"`
}

type ExclusionObservation struct {
}

type ExclusionParameters struct {

	// +kubebuilder:validation:Required
	MatchVariable string `json:"matchVariable" tf:"match_variable"`

	// +kubebuilder:validation:Required
	Selector string `json:"selector" tf:"selector"`

	// +kubebuilder:validation:Required
	SelectorMatchOperator string `json:"selectorMatchOperator" tf:"selector_match_operator"`
}

type ManagedRuleSetObservation struct {
}

type ManagedRuleSetParameters struct {

	// +kubebuilder:validation:Optional
	RuleGroupOverride []RuleGroupOverrideParameters `json:"ruleGroupOverride,omitempty" tf:"rule_group_override"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`

	// +kubebuilder:validation:Required
	Version string `json:"version" tf:"version"`
}

type ManagedRulesObservation struct {
}

type ManagedRulesParameters struct {

	// +kubebuilder:validation:Optional
	Exclusion []ExclusionParameters `json:"exclusion,omitempty" tf:"exclusion"`

	// +kubebuilder:validation:Required
	ManagedRuleSet []ManagedRuleSetParameters `json:"managedRuleSet" tf:"managed_rule_set"`
}

type MatchConditionsObservation struct {
}

type MatchConditionsParameters struct {

	// +kubebuilder:validation:Required
	MatchValues []string `json:"matchValues" tf:"match_values"`

	// +kubebuilder:validation:Required
	MatchVariables []MatchVariablesParameters `json:"matchVariables" tf:"match_variables"`

	// +kubebuilder:validation:Optional
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition"`

	// +kubebuilder:validation:Required
	Operator string `json:"operator" tf:"operator"`

	// +kubebuilder:validation:Optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type MatchVariablesObservation struct {
}

type MatchVariablesParameters struct {

	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector"`

	// +kubebuilder:validation:Required
	VariableName string `json:"variableName" tf:"variable_name"`
}

type PolicySettingsObservation struct {
}

type PolicySettingsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Optional
	FileUploadLimitInMb *int64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb"`

	// +kubebuilder:validation:Optional
	MaxRequestBodySizeInKb *int64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode"`

	// +kubebuilder:validation:Optional
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check"`
}

type RuleGroupOverrideObservation struct {
}

type RuleGroupOverrideParameters struct {

	// +kubebuilder:validation:Required
	DisabledRules []string `json:"disabledRules" tf:"disabled_rules"`

	// +kubebuilder:validation:Required
	RuleGroupName string `json:"ruleGroupName" tf:"rule_group_name"`
}

type WebApplicationFirewallPolicyObservation struct {
	HTTPListenerIds []string `json:"httpListenerIds,omitempty" tf:"http_listener_ids"`

	PathBasedRuleIds []string `json:"pathBasedRuleIds,omitempty" tf:"path_based_rule_ids"`
}

type WebApplicationFirewallPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CustomRules []CustomRulesParameters `json:"customRules,omitempty" tf:"custom_rules"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	ManagedRules []ManagedRulesParameters `json:"managedRules" tf:"managed_rules"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PolicySettings []PolicySettingsParameters `json:"policySettings,omitempty" tf:"policy_settings"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// WebApplicationFirewallPolicySpec defines the desired state of WebApplicationFirewallPolicy
type WebApplicationFirewallPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WebApplicationFirewallPolicyParameters `json:"forProvider"`
}

// WebApplicationFirewallPolicyStatus defines the observed state of WebApplicationFirewallPolicy.
type WebApplicationFirewallPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WebApplicationFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationFirewallPolicy is the Schema for the WebApplicationFirewallPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WebApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebApplicationFirewallPolicySpec   `json:"spec"`
	Status            WebApplicationFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationFirewallPolicyList contains a list of WebApplicationFirewallPolicys
type WebApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApplicationFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	WebApplicationFirewallPolicyKind             = "WebApplicationFirewallPolicy"
	WebApplicationFirewallPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: WebApplicationFirewallPolicyKind}.String()
	WebApplicationFirewallPolicyKindAPIVersion   = WebApplicationFirewallPolicyKind + "." + GroupVersion.String()
	WebApplicationFirewallPolicyGroupVersionKind = GroupVersion.WithKind(WebApplicationFirewallPolicyKind)
)

func init() {
	SchemeBuilder.Register(&WebApplicationFirewallPolicy{}, &WebApplicationFirewallPolicyList{})
}
