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

type AppServiceCertificateOrderObservation struct {
	AppServiceCertificateNotRenewableReasons []string `json:"appServiceCertificateNotRenewableReasons,omitempty" tf:"app_service_certificate_not_renewable_reasons"`

	Certificates []CertificatesObservation `json:"certificates,omitempty" tf:"certificates"`

	DomainVerificationToken string `json:"domainVerificationToken,omitempty" tf:"domain_verification_token"`

	ExpirationTime string `json:"expirationTime,omitempty" tf:"expiration_time"`

	IntermediateThumbprint string `json:"intermediateThumbprint,omitempty" tf:"intermediate_thumbprint"`

	IsPrivateKeyExternal bool `json:"isPrivateKeyExternal,omitempty" tf:"is_private_key_external"`

	RootThumbprint string `json:"rootThumbprint,omitempty" tf:"root_thumbprint"`

	SignedCertificateThumbprint string `json:"signedCertificateThumbprint,omitempty" tf:"signed_certificate_thumbprint"`

	Status string `json:"status,omitempty" tf:"status"`
}

type AppServiceCertificateOrderParameters struct {

	// +kubebuilder:validation:Optional
	AutoRenew *bool `json:"autoRenew,omitempty" tf:"auto_renew"`

	// +kubebuilder:validation:Optional
	Csr *string `json:"csr,omitempty" tf:"csr"`

	// +kubebuilder:validation:Optional
	DistinguishedName *string `json:"distinguishedName,omitempty" tf:"distinguished_name"`

	// +kubebuilder:validation:Optional
	KeySize *int64 `json:"keySize,omitempty" tf:"key_size"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	ProductType *string `json:"productType,omitempty" tf:"product_type"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	ValidityInYears *int64 `json:"validityInYears,omitempty" tf:"validity_in_years"`
}

type CertificatesObservation struct {
	CertificateName string `json:"certificateName,omitempty" tf:"certificate_name"`

	KeyVaultID string `json:"keyVaultId,omitempty" tf:"key_vault_id"`

	KeyVaultSecretName string `json:"keyVaultSecretName,omitempty" tf:"key_vault_secret_name"`

	ProvisioningState string `json:"provisioningState,omitempty" tf:"provisioning_state"`
}

type CertificatesParameters struct {
}

// AppServiceCertificateOrderSpec defines the desired state of AppServiceCertificateOrder
type AppServiceCertificateOrderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppServiceCertificateOrderParameters `json:"forProvider"`
}

// AppServiceCertificateOrderStatus defines the observed state of AppServiceCertificateOrder.
type AppServiceCertificateOrderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppServiceCertificateOrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceCertificateOrder is the Schema for the AppServiceCertificateOrders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppServiceCertificateOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceCertificateOrderSpec   `json:"spec"`
	Status            AppServiceCertificateOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceCertificateOrderList contains a list of AppServiceCertificateOrders
type AppServiceCertificateOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServiceCertificateOrder `json:"items"`
}

// Repository type metadata.
var (
	AppServiceCertificateOrderKind             = "AppServiceCertificateOrder"
	AppServiceCertificateOrderGroupKind        = schema.GroupKind{Group: Group, Kind: AppServiceCertificateOrderKind}.String()
	AppServiceCertificateOrderKindAPIVersion   = AppServiceCertificateOrderKind + "." + GroupVersion.String()
	AppServiceCertificateOrderGroupVersionKind = GroupVersion.WithKind(AppServiceCertificateOrderKind)
)

func init() {
	SchemeBuilder.Register(&AppServiceCertificateOrder{}, &AppServiceCertificateOrderList{})
}
