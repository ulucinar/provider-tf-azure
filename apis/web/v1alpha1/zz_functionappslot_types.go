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

type FunctionAppSlotAuthSettingsActiveDirectoryObservation struct {
}

type FunctionAppSlotAuthSettingsActiveDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	AllowedAudiences []*string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`
}

type FunctionAppSlotAuthSettingsFacebookObservation struct {
}

type FunctionAppSlotAuthSettingsFacebookParameters struct {

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Required
	AppSecretSecretRef v1.SecretKeySelector `json:"appSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type FunctionAppSlotAuthSettingsGoogleObservation struct {
}

type FunctionAppSlotAuthSettingsGoogleParameters struct {

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type FunctionAppSlotAuthSettingsMicrosoftObservation struct {
}

type FunctionAppSlotAuthSettingsMicrosoftParameters struct {

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type FunctionAppSlotAuthSettingsObservation struct {
}

type FunctionAppSlotAuthSettingsParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectory []FunctionAppSlotAuthSettingsActiveDirectoryParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalLoginParams map[string]*string `json:"additionalLoginParams,omitempty" tf:"additional_login_params,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedExternalRedirectUrls []*string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultProvider *string `json:"defaultProvider,omitempty" tf:"default_provider,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Facebook []FunctionAppSlotAuthSettingsFacebookParameters `json:"facebook,omitempty" tf:"facebook,omitempty"`

	// +kubebuilder:validation:Optional
	Google []FunctionAppSlotAuthSettingsGoogleParameters `json:"google,omitempty" tf:"google,omitempty"`

	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// +kubebuilder:validation:Optional
	Microsoft []FunctionAppSlotAuthSettingsMicrosoftParameters `json:"microsoft,omitempty" tf:"microsoft,omitempty"`

	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`

	// +kubebuilder:validation:Optional
	TokenRefreshExtensionHours *float64 `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours,omitempty"`

	// +kubebuilder:validation:Optional
	TokenStoreEnabled *bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Twitter []FunctionAppSlotAuthSettingsTwitterParameters `json:"twitter,omitempty" tf:"twitter,omitempty"`

	// +kubebuilder:validation:Optional
	UnauthenticatedClientAction *string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action,omitempty"`
}

type FunctionAppSlotAuthSettingsTwitterObservation struct {
}

type FunctionAppSlotAuthSettingsTwitterParameters struct {

	// +kubebuilder:validation:Required
	ConsumerKey *string `json:"consumerKey" tf:"consumer_key,omitempty"`

	// +kubebuilder:validation:Required
	ConsumerSecretSecretRef v1.SecretKeySelector `json:"consumerSecretSecretRef" tf:"-"`
}

type FunctionAppSlotConnectionStringObservation struct {
}

type FunctionAppSlotConnectionStringParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type FunctionAppSlotIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type FunctionAppSlotIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FunctionAppSlotObservation struct {
	DefaultHostName *string `json:"defaultHostname,omitempty" tf:"default_hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	OutboundIPAddresses *string `json:"outboundIpAddresses,omitempty" tf:"outbound_ip_addresses,omitempty"`

	PossibleOutboundIPAddresses *string `json:"possibleOutboundIpAddresses,omitempty" tf:"possible_outbound_ip_addresses,omitempty"`

	SiteCredential []FunctionAppSlotSiteCredentialObservation `json:"siteCredential,omitempty" tf:"site_credential,omitempty"`
}

type FunctionAppSlotParameters struct {

	// +kubebuilder:validation:Required
	AppServicePlanID *string `json:"appServicePlanId" tf:"app_service_plan_id,omitempty"`

	// +kubebuilder:validation:Optional
	AppSettings map[string]*string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`

	// +kubebuilder:validation:Optional
	AuthSettings []FunctionAppSlotAuthSettingsParameters `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`

	// +kubebuilder:validation:Optional
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionString []FunctionAppSlotConnectionStringParameters `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// +kubebuilder:validation:Optional
	DailyMemoryTimeQuota *int64 `json:"dailyMemoryTimeQuota,omitempty" tf:"daily_memory_time_quota,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBuiltinLogging *bool `json:"enableBuiltinLogging,omitempty" tf:"enable_builtin_logging,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	FunctionAppName *string `json:"functionAppName" tf:"function_app_name,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []FunctionAppSlotIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SiteConfig []FunctionAppSlotSiteConfigParameters `json:"siteConfig,omitempty" tf:"site_config,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountAccessKeySecretRef v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageAccountName *string `json:"storageAccountName" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionAppSlotSiteConfigCorsObservation struct {
}

type FunctionAppSlotSiteConfigCorsParameters struct {

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Optional
	SupportCredentials *bool `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type FunctionAppSlotSiteConfigIPRestrictionHeadersObservation struct {
}

type FunctionAppSlotSiteConfigIPRestrictionHeadersParameters struct {

	// +kubebuilder:validation:Optional
	XAzureFdid []*string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid,omitempty"`

	// +kubebuilder:validation:Optional
	XFdHealthProbe []*string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe,omitempty"`

	// +kubebuilder:validation:Optional
	XForwardedFor []*string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for,omitempty"`

	// +kubebuilder:validation:Optional
	XForwardedHost []*string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host,omitempty"`
}

type FunctionAppSlotSiteConfigIPRestrictionObservation struct {
}

type FunctionAppSlotSiteConfigIPRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []FunctionAppSlotSiteConfigIPRestrictionHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDRef *v1.Reference `json:"virtualNetworkSubnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDSelector *v1.Selector `json:"virtualNetworkSubnetIdSelector,omitempty" tf:"-"`
}

type FunctionAppSlotSiteConfigObservation struct {
}

type FunctionAppSlotSiteConfigParameters struct {

	// +kubebuilder:validation:Optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`

	// +kubebuilder:validation:Optional
	AppScaleLimit *int64 `json:"appScaleLimit,omitempty" tf:"app_scale_limit,omitempty"`

	// +kubebuilder:validation:Optional
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty" tf:"auto_swap_slot_name,omitempty"`

	// +kubebuilder:validation:Optional
	Cors []FunctionAppSlotSiteConfigCorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// +kubebuilder:validation:Optional
	DotnetFrameworkVersion *string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticInstanceMinimum *int64 `json:"elasticInstanceMinimum,omitempty" tf:"elastic_instance_minimum,omitempty"`

	// +kubebuilder:validation:Optional
	FtpsState *string `json:"ftpsState,omitempty" tf:"ftps_state,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// +kubebuilder:validation:Optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IPRestriction []FunctionAppSlotSiteConfigIPRestrictionParameters `json:"ipRestriction,omitempty" tf:"ip_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	JavaVersion *string `json:"javaVersion,omitempty" tf:"java_version,omitempty"`

	// +kubebuilder:validation:Optional
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`

	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// +kubebuilder:validation:Optional
	PreWarmedInstanceCount *int64 `json:"preWarmedInstanceCount,omitempty" tf:"pre_warmed_instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	RuntimeScaleMonitoringEnabled *bool `json:"runtimeScaleMonitoringEnabled,omitempty" tf:"runtime_scale_monitoring_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ScmIPRestriction []FunctionAppSlotSiteConfigScmIPRestrictionParameters `json:"scmIpRestriction,omitempty" tf:"scm_ip_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	ScmType *string `json:"scmType,omitempty" tf:"scm_type,omitempty"`

	// +kubebuilder:validation:Optional
	ScmUseMainIPRestriction *bool `json:"scmUseMainIpRestriction,omitempty" tf:"scm_use_main_ip_restriction,omitempty"`

	// +kubebuilder:validation:Optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`

	// +kubebuilder:validation:Optional
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty" tf:"vnet_route_all_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled,omitempty"`
}

type FunctionAppSlotSiteConfigScmIPRestrictionHeadersObservation struct {
}

type FunctionAppSlotSiteConfigScmIPRestrictionHeadersParameters struct {

	// +kubebuilder:validation:Optional
	XAzureFdid []*string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid,omitempty"`

	// +kubebuilder:validation:Optional
	XFdHealthProbe []*string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe,omitempty"`

	// +kubebuilder:validation:Optional
	XForwardedFor []*string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for,omitempty"`

	// +kubebuilder:validation:Optional
	XForwardedHost []*string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host,omitempty"`
}

type FunctionAppSlotSiteConfigScmIPRestrictionObservation struct {
}

type FunctionAppSlotSiteConfigScmIPRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []FunctionAppSlotSiteConfigScmIPRestrictionHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDRef *v1.Reference `json:"virtualNetworkSubnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDSelector *v1.Selector `json:"virtualNetworkSubnetIdSelector,omitempty" tf:"-"`
}

type FunctionAppSlotSiteCredentialObservation struct {
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type FunctionAppSlotSiteCredentialParameters struct {
}

// FunctionAppSlotSpec defines the desired state of FunctionAppSlot
type FunctionAppSlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionAppSlotParameters `json:"forProvider"`
}

// FunctionAppSlotStatus defines the observed state of FunctionAppSlot.
type FunctionAppSlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionAppSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppSlot is the Schema for the FunctionAppSlots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FunctionAppSlot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppSlotSpec   `json:"spec"`
	Status            FunctionAppSlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppSlotList contains a list of FunctionAppSlots
type FunctionAppSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionAppSlot `json:"items"`
}

// Repository type metadata.
var (
	FunctionAppSlot_Kind             = "FunctionAppSlot"
	FunctionAppSlot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionAppSlot_Kind}.String()
	FunctionAppSlot_KindAPIVersion   = FunctionAppSlot_Kind + "." + CRDGroupVersion.String()
	FunctionAppSlot_GroupVersionKind = CRDGroupVersion.WithKind(FunctionAppSlot_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppSlot{}, &FunctionAppSlotList{})
}
