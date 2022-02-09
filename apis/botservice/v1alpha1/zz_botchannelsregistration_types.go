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

type BotChannelsRegistrationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BotChannelsRegistrationParameters struct {

	// +kubebuilder:validation:Optional
	CmkKeyVaultURL *string `json:"cmkKeyVaultUrl,omitempty" tf:"cmk_key_vault_url,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeveloperAppInsightsAPIKeySecretRef *v1.SecretKeySelector `json:"developerAppInsightsApiKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// +kubebuilder:validation:Optional
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// +kubebuilder:validation:Optional
	IsolatedNetworkEnabled *bool `json:"isolatedNetworkEnabled,omitempty" tf:"isolated_network_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	MicrosoftAppID *string `json:"microsoftAppId" tf:"microsoft_app_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BotChannelsRegistrationSpec defines the desired state of BotChannelsRegistration
type BotChannelsRegistrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelsRegistrationParameters `json:"forProvider"`
}

// BotChannelsRegistrationStatus defines the observed state of BotChannelsRegistration.
type BotChannelsRegistrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelsRegistrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelsRegistration is the Schema for the BotChannelsRegistrations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type BotChannelsRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelsRegistrationSpec   `json:"spec"`
	Status            BotChannelsRegistrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelsRegistrationList contains a list of BotChannelsRegistrations
type BotChannelsRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelsRegistration `json:"items"`
}

// Repository type metadata.
var (
	BotChannelsRegistration_Kind             = "BotChannelsRegistration"
	BotChannelsRegistration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelsRegistration_Kind}.String()
	BotChannelsRegistration_KindAPIVersion   = BotChannelsRegistration_Kind + "." + CRDGroupVersion.String()
	BotChannelsRegistration_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelsRegistration_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelsRegistration{}, &BotChannelsRegistrationList{})
}
