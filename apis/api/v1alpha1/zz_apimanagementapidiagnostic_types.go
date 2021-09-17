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

type ApiManagementApiDiagnosticObservation struct {
}

type ApiManagementApiDiagnosticParameters struct {

	// +kubebuilder:validation:Required
	APIManagementLoggerID string `json:"apiManagementLoggerId" tf:"api_management_logger_id"`

	// +kubebuilder:validation:Required
	APIManagementName string `json:"apiManagementName" tf:"api_management_name"`

	// +kubebuilder:validation:Required
	APIName string `json:"apiName" tf:"api_name"`

	// +kubebuilder:validation:Optional
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors"`

	// +kubebuilder:validation:Optional
	BackendRequest []BackendRequestParameters `json:"backendRequest,omitempty" tf:"backend_request"`

	// +kubebuilder:validation:Optional
	BackendResponse []BackendResponseParameters `json:"backendResponse,omitempty" tf:"backend_response"`

	// +kubebuilder:validation:Optional
	FrontendRequest []FrontendRequestParameters `json:"frontendRequest,omitempty" tf:"frontend_request"`

	// +kubebuilder:validation:Optional
	FrontendResponse []FrontendResponseParameters `json:"frontendResponse,omitempty" tf:"frontend_response"`

	// +kubebuilder:validation:Optional
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol"`

	// +kubebuilder:validation:Required
	Identifier string `json:"identifier" tf:"identifier"`

	// +kubebuilder:validation:Optional
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip"`

	// +kubebuilder:validation:Optional
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage"`

	// +kubebuilder:validation:Optional
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity"`
}

type BackendRequestObservation struct {
}

type BackendRequestParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	// +kubebuilder:validation:Optional
	DataMasking []DataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	// +kubebuilder:validation:Optional
	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type BackendResponseDataMaskingObservation struct {
}

type BackendResponseDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []DataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// +kubebuilder:validation:Optional
	QueryParams []DataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type BackendResponseObservation struct {
}

type BackendResponseParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	// +kubebuilder:validation:Optional
	DataMasking []BackendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	// +kubebuilder:validation:Optional
	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type DataMaskingHeadersObservation struct {
}

type DataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type DataMaskingObservation struct {
}

type DataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers"`

	// +kubebuilder:validation:Optional
	QueryParams []QueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type DataMaskingQueryParamsObservation struct {
}

type DataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FrontendRequestDataMaskingHeadersObservation struct {
}

type FrontendRequestDataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FrontendRequestDataMaskingObservation struct {
}

type FrontendRequestDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []FrontendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// +kubebuilder:validation:Optional
	QueryParams []FrontendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type FrontendRequestDataMaskingQueryParamsObservation struct {
}

type FrontendRequestDataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FrontendRequestObservation struct {
}

type FrontendRequestParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	// +kubebuilder:validation:Optional
	DataMasking []FrontendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	// +kubebuilder:validation:Optional
	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type FrontendResponseDataMaskingHeadersObservation struct {
}

type FrontendResponseDataMaskingHeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FrontendResponseDataMaskingObservation struct {
}

type FrontendResponseDataMaskingParameters struct {

	// +kubebuilder:validation:Optional
	Headers []FrontendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// +kubebuilder:validation:Optional
	QueryParams []FrontendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params"`
}

type FrontendResponseDataMaskingQueryParamsObservation struct {
}

type FrontendResponseDataMaskingQueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FrontendResponseObservation struct {
}

type FrontendResponseParameters struct {

	// +kubebuilder:validation:Optional
	BodyBytes *int64 `json:"bodyBytes,omitempty" tf:"body_bytes"`

	// +kubebuilder:validation:Optional
	DataMasking []FrontendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking"`

	// +kubebuilder:validation:Optional
	HeadersToLog []string `json:"headersToLog,omitempty" tf:"headers_to_log"`
}

type HeadersObservation struct {
}

type HeadersParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type QueryParamsObservation struct {
}

type QueryParamsParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

// ApiManagementApiDiagnosticSpec defines the desired state of ApiManagementApiDiagnostic
type ApiManagementApiDiagnosticSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementApiDiagnosticParameters `json:"forProvider"`
}

// ApiManagementApiDiagnosticStatus defines the observed state of ApiManagementApiDiagnostic.
type ApiManagementApiDiagnosticStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementApiDiagnosticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiDiagnostic is the Schema for the ApiManagementApiDiagnostics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementApiDiagnostic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiDiagnosticSpec   `json:"spec"`
	Status            ApiManagementApiDiagnosticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiDiagnosticList contains a list of ApiManagementApiDiagnostics
type ApiManagementApiDiagnosticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementApiDiagnostic `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementApiDiagnosticKind             = "ApiManagementApiDiagnostic"
	ApiManagementApiDiagnosticGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementApiDiagnosticKind}.String()
	ApiManagementApiDiagnosticKindAPIVersion   = ApiManagementApiDiagnosticKind + "." + GroupVersion.String()
	ApiManagementApiDiagnosticGroupVersionKind = GroupVersion.WithKind(ApiManagementApiDiagnosticKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementApiDiagnostic{}, &ApiManagementApiDiagnosticList{})
}
