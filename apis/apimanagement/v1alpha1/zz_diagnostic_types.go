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

type BackendRequestDataMaskingHeadersObservation struct {
}

type BackendRequestDataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendRequestDataMaskingObservation struct {
}

type BackendRequestDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []BackendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	QueryParams []BackendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendRequestDataMaskingQueryParamsObservation struct {
}

type BackendRequestDataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendResponseDataMaskingHeadersObservation struct {
}

type BackendResponseDataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendResponseDataMaskingQueryParamsObservation struct {
}

type BackendResponseDataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticBackendRequestObservation struct {
}

type DiagnosticBackendRequestParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	DataMasking []BackendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendResponseDataMaskingObservation struct {
}

type DiagnosticBackendResponseDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []BackendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	QueryParams []BackendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticBackendResponseObservation struct {
}

type DiagnosticBackendResponseParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	DataMasking []DiagnosticBackendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingHeadersObservation struct {
}

type DiagnosticFrontendRequestDataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingObservation struct {
}

type DiagnosticFrontendRequestDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []DiagnosticFrontendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	QueryParams []DiagnosticFrontendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingQueryParamsObservation struct {
}

type DiagnosticFrontendRequestDataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestObservation struct {
}

type DiagnosticFrontendRequestParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	DataMasking []DiagnosticFrontendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingHeadersObservation struct {
}

type DiagnosticFrontendResponseDataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingObservation struct {
}

type DiagnosticFrontendResponseDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []DiagnosticFrontendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	QueryParams []DiagnosticFrontendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingQueryParamsObservation struct {
}

type DiagnosticFrontendResponseDataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseObservation struct {
}

type DiagnosticFrontendResponseParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	DataMasking []DiagnosticFrontendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiagnosticParameters struct {

	// +kubebuilder:validation:Required
	APIManagementLoggerID *string `json:"apiManagementLoggerId" tf:"api_management_logger_id,omitempty"`

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// +kubebuilder:validation:Optional
	BackendRequest []DiagnosticBackendRequestParameters `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// +kubebuilder:validation:Optional
	BackendResponse []DiagnosticBackendResponseParameters `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	FrontendRequest []DiagnosticFrontendRequestParameters `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// +kubebuilder:validation:Optional
	FrontendResponse []DiagnosticFrontendResponseParameters `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// +kubebuilder:validation:Required
	Identifier *string `json:"identifier" tf:"identifier,omitempty"`

	// +kubebuilder:validation:Optional
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// +kubebuilder:validation:Optional
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

// DiagnosticSpec defines the desired state of Diagnostic
type DiagnosticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiagnosticParameters `json:"forProvider"`
}

// DiagnosticStatus defines the observed state of Diagnostic.
type DiagnosticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiagnosticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Diagnostic is the Schema for the Diagnostics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Diagnostic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiagnosticSpec   `json:"spec"`
	Status            DiagnosticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiagnosticList contains a list of Diagnostics
type DiagnosticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Diagnostic `json:"items"`
}

// Repository type metadata.
var (
	Diagnostic_Kind             = "Diagnostic"
	Diagnostic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Diagnostic_Kind}.String()
	Diagnostic_KindAPIVersion   = Diagnostic_Kind + "." + CRDGroupVersion.String()
	Diagnostic_GroupVersionKind = CRDGroupVersion.WithKind(Diagnostic_Kind)
)

func init() {
	SchemeBuilder.Register(&Diagnostic{}, &DiagnosticList{})
}
