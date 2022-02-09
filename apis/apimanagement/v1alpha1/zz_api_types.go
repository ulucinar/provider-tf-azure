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

type APIObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsCurrent *bool `json:"isCurrent,omitempty" tf:"is_current,omitempty"`

	IsOnline *bool `json:"isOnline,omitempty" tf:"is_online,omitempty"`
}

type APIParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Oauth2Authorization []Oauth2AuthorizationParameters `json:"oauth2Authorization,omitempty" tf:"oauth2_authorization,omitempty"`

	// +kubebuilder:validation:Optional
	OpenIDAuthentication []OpenIDAuthenticationParameters `json:"openidAuthentication,omitempty" tf:"openid_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Revision *string `json:"revision" tf:"revision,omitempty"`

	// +kubebuilder:validation:Optional
	RevisionDescription *string `json:"revisionDescription,omitempty" tf:"revision_description,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// +kubebuilder:validation:Optional
	SoapPassThrough *bool `json:"soapPassThrough,omitempty" tf:"soap_pass_through,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAPIID *string `json:"sourceApiId,omitempty" tf:"source_api_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubscriptionKeyParameterNames []SubscriptionKeyParameterNamesParameters `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names,omitempty"`

	// +kubebuilder:validation:Optional
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Optional
	VersionDescription *string `json:"versionDescription,omitempty" tf:"version_description,omitempty"`

	// +kubebuilder:validation:Optional
	VersionSetID *string `json:"versionSetId,omitempty" tf:"version_set_id,omitempty"`
}

type ImportObservation struct {
}

type ImportParameters struct {

	// +kubebuilder:validation:Required
	ContentFormat *string `json:"contentFormat" tf:"content_format,omitempty"`

	// +kubebuilder:validation:Required
	ContentValue *string `json:"contentValue" tf:"content_value,omitempty"`

	// +kubebuilder:validation:Optional
	WsdlSelector []WsdlSelectorParameters `json:"wsdlSelector,omitempty" tf:"wsdl_selector,omitempty"`
}

type Oauth2AuthorizationObservation struct {
}

type Oauth2AuthorizationParameters struct {

	// +kubebuilder:validation:Required
	AuthorizationServerName *string `json:"authorizationServerName" tf:"authorization_server_name,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type OpenIDAuthenticationObservation struct {
}

type OpenIDAuthenticationParameters struct {

	// +kubebuilder:validation:Optional
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// +kubebuilder:validation:Required
	OpenIDProviderName *string `json:"openidProviderName" tf:"openid_provider_name,omitempty"`
}

type SubscriptionKeyParameterNamesObservation struct {
}

type SubscriptionKeyParameterNamesParameters struct {

	// +kubebuilder:validation:Required
	Header *string `json:"header" tf:"header,omitempty"`

	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`
}

type WsdlSelectorObservation struct {
}

type WsdlSelectorParameters struct {

	// +kubebuilder:validation:Required
	EndpointName *string `json:"endpointName" tf:"endpoint_name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// APISpec defines the desired state of API
type APISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIParameters `json:"forProvider"`
}

// APIStatus defines the observed state of API.
type APIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// API is the Schema for the APIs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APISpec   `json:"spec"`
	Status            APIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIList contains a list of APIs
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

// Repository type metadata.
var (
	API_Kind             = "API"
	API_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: API_Kind}.String()
	API_KindAPIVersion   = API_Kind + "." + CRDGroupVersion.String()
	API_GroupVersionKind = CRDGroupVersion.WithKind(API_Kind)
)

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
