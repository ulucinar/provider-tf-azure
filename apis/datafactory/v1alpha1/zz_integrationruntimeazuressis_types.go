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

type CatalogInfoObservation struct {
}

type CatalogInfoParameters struct {

	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// +kubebuilder:validation:Optional
	AdministratorPasswordSecretRef *v1.SecretKeySelector `json:"administratorPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DualStandbyPairName *string `json:"dualStandbyPairName,omitempty" tf:"dual_standby_pair_name,omitempty"`

	// +kubebuilder:validation:Optional
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`

	// +kubebuilder:validation:Required
	ServerEndpoint *string `json:"serverEndpoint" tf:"server_endpoint,omitempty"`
}

type CommandKeyObservation struct {
}

type CommandKeyParameters struct {

	// +kubebuilder:validation:Optional
	KeyVaultPassword []KeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TargetName *string `json:"targetName" tf:"target_name,omitempty"`

	// +kubebuilder:validation:Required
	UserName *string `json:"userName" tf:"user_name,omitempty"`
}

type ComponentObservation struct {
}

type ComponentParameters struct {

	// +kubebuilder:validation:Optional
	KeyVaultLicense []KeyVaultLicenseParameters `json:"keyVaultLicense,omitempty" tf:"key_vault_license,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseSecretRef *v1.SecretKeySelector `json:"licenseSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type CustomSetupScriptObservation struct {
}

type CustomSetupScriptParameters struct {

	// +kubebuilder:validation:Required
	BlobContainerURI *string `json:"blobContainerUri" tf:"blob_container_uri,omitempty"`

	// +kubebuilder:validation:Required
	SASTokenSecretRef v1.SecretKeySelector `json:"sasTokenSecretRef" tf:"-"`
}

type ExpressCustomSetupObservation struct {
}

type ExpressCustomSetupParameters struct {

	// +kubebuilder:validation:Optional
	CommandKey []CommandKeyParameters `json:"commandKey,omitempty" tf:"command_key,omitempty"`

	// +kubebuilder:validation:Optional
	Component []ComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// +kubebuilder:validation:Optional
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	PowershellVersion *string `json:"powershellVersion,omitempty" tf:"powershell_version,omitempty"`
}

type IntegrationRuntimeAzureSSISObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IntegrationRuntimeAzureSSISParameters struct {

	// +kubebuilder:validation:Optional
	CatalogInfo []CatalogInfoParameters `json:"catalogInfo,omitempty" tf:"catalog_info,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSetupScript []CustomSetupScriptParameters `json:"customSetupScript,omitempty" tf:"custom_setup_script,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// +kubebuilder:validation:Optional
	ExpressCustomSetup []ExpressCustomSetupParameters `json:"expressCustomSetup,omitempty" tf:"express_custom_setup,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaxParallelExecutionsPerNode *int64 `json:"maxParallelExecutionsPerNode,omitempty" tf:"max_parallel_executions_per_node,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NodeSize *string `json:"nodeSize" tf:"node_size,omitempty"`

	// +kubebuilder:validation:Optional
	NumberOfNodes *int64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// +kubebuilder:validation:Optional
	PackageStore []PackageStoreParameters `json:"packageStore,omitempty" tf:"package_store,omitempty"`

	// +kubebuilder:validation:Optional
	Proxy []ProxyParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VnetIntegration []VnetIntegrationParameters `json:"vnetIntegration,omitempty" tf:"vnet_integration,omitempty"`
}

type KeyVaultLicenseObservation struct {
}

type KeyVaultLicenseParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecretVersion *string `json:"secretVersion,omitempty" tf:"secret_version,omitempty"`
}

type KeyVaultPasswordObservation struct {
}

type KeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecretVersion *string `json:"secretVersion,omitempty" tf:"secret_version,omitempty"`
}

type PackageStoreObservation struct {
}

type PackageStoreParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ProxyObservation struct {
}

type ProxyParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	SelfHostedIntegrationRuntimeName *string `json:"selfHostedIntegrationRuntimeName" tf:"self_hosted_integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	StagingStorageLinkedServiceName *string `json:"stagingStorageLinkedServiceName" tf:"staging_storage_linked_service_name,omitempty"`
}

type VnetIntegrationObservation struct {
}

type VnetIntegrationParameters struct {

	// +kubebuilder:validation:Optional
	PublicIps []*string `json:"publicIps,omitempty" tf:"public_ips,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +kubebuilder:validation:Optional
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetNameRef *v1.Reference `json:"subnetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetNameSelector *v1.Selector `json:"subnetNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VnetID *string `json:"vnetId,omitempty" tf:"vnet_id,omitempty"`
}

// IntegrationRuntimeAzureSSISSpec defines the desired state of IntegrationRuntimeAzureSSIS
type IntegrationRuntimeAzureSSISSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationRuntimeAzureSSISParameters `json:"forProvider"`
}

// IntegrationRuntimeAzureSSISStatus defines the observed state of IntegrationRuntimeAzureSSIS.
type IntegrationRuntimeAzureSSISStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationRuntimeAzureSSISObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationRuntimeAzureSSIS is the Schema for the IntegrationRuntimeAzureSSISs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IntegrationRuntimeAzureSSIS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationRuntimeAzureSSISSpec   `json:"spec"`
	Status            IntegrationRuntimeAzureSSISStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationRuntimeAzureSSISList contains a list of IntegrationRuntimeAzureSSISs
type IntegrationRuntimeAzureSSISList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationRuntimeAzureSSIS `json:"items"`
}

// Repository type metadata.
var (
	IntegrationRuntimeAzureSSIS_Kind             = "IntegrationRuntimeAzureSSIS"
	IntegrationRuntimeAzureSSIS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationRuntimeAzureSSIS_Kind}.String()
	IntegrationRuntimeAzureSSIS_KindAPIVersion   = IntegrationRuntimeAzureSSIS_Kind + "." + CRDGroupVersion.String()
	IntegrationRuntimeAzureSSIS_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationRuntimeAzureSSIS_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationRuntimeAzureSSIS{}, &IntegrationRuntimeAzureSSISList{})
}
