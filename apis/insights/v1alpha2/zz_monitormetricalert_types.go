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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Required
	ActionGroupID *string `json:"actionGroupId" tf:"action_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	WebhookProperties map[string]*string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type ApplicationInsightsWebTestLocationAvailabilityCriteriaObservation struct {
}

type ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters struct {

	// +kubebuilder:validation:Required
	ComponentID *string `json:"componentId" tf:"component_id,omitempty"`

	// +kubebuilder:validation:Required
	FailedLocationCount *int64 `json:"failedLocationCount" tf:"failed_location_count,omitempty"`

	// +kubebuilder:validation:Required
	WebTestID *string `json:"webTestId" tf:"web_test_id,omitempty"`
}

type CriteriaObservation struct {
}

type CriteriaParameters struct {

	// +kubebuilder:validation:Required
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// +kubebuilder:validation:Optional
	Dimension []DimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Required
	MetricNamespace *string `json:"metricNamespace" tf:"metric_namespace,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	SkipMetricValidation *bool `json:"skipMetricValidation,omitempty" tf:"skip_metric_validation,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type DimensionObservation struct {
}

type DimensionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type DynamicCriteriaDimensionObservation struct {
}

type DynamicCriteriaDimensionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type DynamicCriteriaObservation struct {
}

type DynamicCriteriaParameters struct {

	// +kubebuilder:validation:Required
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// +kubebuilder:validation:Required
	AlertSensitivity *string `json:"alertSensitivity" tf:"alert_sensitivity,omitempty"`

	// +kubebuilder:validation:Optional
	Dimension []DynamicCriteriaDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// +kubebuilder:validation:Optional
	EvaluationFailureCount *int64 `json:"evaluationFailureCount,omitempty" tf:"evaluation_failure_count,omitempty"`

	// +kubebuilder:validation:Optional
	EvaluationTotalCount *int64 `json:"evaluationTotalCount,omitempty" tf:"evaluation_total_count,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreDataBefore *string `json:"ignoreDataBefore,omitempty" tf:"ignore_data_before,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Required
	MetricNamespace *string `json:"metricNamespace" tf:"metric_namespace,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	SkipMetricValidation *bool `json:"skipMetricValidation,omitempty" tf:"skip_metric_validation,omitempty"`
}

type MonitorMetricAlertObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorMetricAlertParameters struct {

	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationInsightsWebTestLocationAvailabilityCriteria []ApplicationInsightsWebTestLocationAvailabilityCriteriaParameters `json:"applicationInsightsWebTestLocationAvailabilityCriteria,omitempty" tf:"application_insights_web_test_location_availability_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	AutoMitigate *bool `json:"autoMitigate,omitempty" tf:"auto_mitigate,omitempty"`

	// +kubebuilder:validation:Optional
	Criteria []CriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicCriteria []DynamicCriteriaParameters `json:"dynamicCriteria,omitempty" tf:"dynamic_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`

	// +kubebuilder:validation:Optional
	Severity *int64 `json:"severity,omitempty" tf:"severity,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The location of the target pluginsdk. Required when using subscription, resource group scope or multiple scopes.
	// +kubebuilder:validation:Optional
	TargetResourceLocation *string `json:"targetResourceLocation,omitempty" tf:"target_resource_location,omitempty"`

	// The resource type (e.g. Microsoft.Compute/virtualMachines) of the target pluginsdk. Required when using subscription, resource group scope or multiple scopes.
	// +kubebuilder:validation:Optional
	TargetResourceType *string `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	WindowSize *string `json:"windowSize,omitempty" tf:"window_size,omitempty"`
}

// MonitorMetricAlertSpec defines the desired state of MonitorMetricAlert
type MonitorMetricAlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorMetricAlertParameters `json:"forProvider"`
}

// MonitorMetricAlertStatus defines the observed state of MonitorMetricAlert.
type MonitorMetricAlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorMetricAlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorMetricAlert is the Schema for the MonitorMetricAlerts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MonitorMetricAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorMetricAlertSpec   `json:"spec"`
	Status            MonitorMetricAlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorMetricAlertList contains a list of MonitorMetricAlerts
type MonitorMetricAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorMetricAlert `json:"items"`
}

// Repository type metadata.
var (
	MonitorMetricAlert_Kind             = "MonitorMetricAlert"
	MonitorMetricAlert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorMetricAlert_Kind}.String()
	MonitorMetricAlert_KindAPIVersion   = MonitorMetricAlert_Kind + "." + CRDGroupVersion.String()
	MonitorMetricAlert_GroupVersionKind = CRDGroupVersion.WithKind(MonitorMetricAlert_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorMetricAlert{}, &MonitorMetricAlertList{})
}
