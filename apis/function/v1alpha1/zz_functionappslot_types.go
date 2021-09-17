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

type AuthSettingsActiveDirectoryObservation struct {
}

type AuthSettingsActiveDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	AllowedAudiences []string `json:"allowedAudiences,omitempty" tf:"allowed_audiences"`

	// +kubebuilder:validation:Required
	ClientID string `json:"clientId" tf:"client_id"`

	// +kubebuilder:validation:Optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret"`
}

type AuthSettingsFacebookObservation struct {
}

type AuthSettingsFacebookParameters struct {

	// +kubebuilder:validation:Required
	AppID string `json:"appId" tf:"app_id"`

	// +kubebuilder:validation:Required
	AppSecret string `json:"appSecret" tf:"app_secret"`

	// +kubebuilder:validation:Optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type AuthSettingsGoogleObservation struct {
}

type AuthSettingsGoogleParameters struct {

	// +kubebuilder:validation:Required
	ClientID string `json:"clientId" tf:"client_id"`

	// +kubebuilder:validation:Required
	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	// +kubebuilder:validation:Optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type AuthSettingsMicrosoftObservation struct {
}

type AuthSettingsMicrosoftParameters struct {

	// +kubebuilder:validation:Required
	ClientID string `json:"clientId" tf:"client_id"`

	// +kubebuilder:validation:Required
	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	// +kubebuilder:validation:Optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type AuthSettingsTwitterObservation struct {
}

type AuthSettingsTwitterParameters struct {

	// +kubebuilder:validation:Required
	ConsumerKey string `json:"consumerKey" tf:"consumer_key"`

	// +kubebuilder:validation:Required
	ConsumerSecret string `json:"consumerSecret" tf:"consumer_secret"`
}

type FunctionAppSlotAuthSettingsObservation struct {
}

type FunctionAppSlotAuthSettingsParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectory []AuthSettingsActiveDirectoryParameters `json:"activeDirectory,omitempty" tf:"active_directory"`

	// +kubebuilder:validation:Optional
	AdditionalLoginParams map[string]string `json:"additionalLoginParams,omitempty" tf:"additional_login_params"`

	// +kubebuilder:validation:Optional
	AllowedExternalRedirectUrls []string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls"`

	// +kubebuilder:validation:Optional
	DefaultProvider *string `json:"defaultProvider,omitempty" tf:"default_provider"`

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Optional
	Facebook []AuthSettingsFacebookParameters `json:"facebook,omitempty" tf:"facebook"`

	// +kubebuilder:validation:Optional
	Google []AuthSettingsGoogleParameters `json:"google,omitempty" tf:"google"`

	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer"`

	// +kubebuilder:validation:Optional
	Microsoft []AuthSettingsMicrosoftParameters `json:"microsoft,omitempty" tf:"microsoft"`

	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version"`

	// +kubebuilder:validation:Optional
	TokenRefreshExtensionHours *float64 `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours"`

	// +kubebuilder:validation:Optional
	TokenStoreEnabled *bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled"`

	// +kubebuilder:validation:Optional
	Twitter []AuthSettingsTwitterParameters `json:"twitter,omitempty" tf:"twitter"`

	// +kubebuilder:validation:Optional
	UnauthenticatedClientAction *string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action"`
}

type FunctionAppSlotConnectionStringObservation struct {
}

type FunctionAppSlotConnectionStringParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type FunctionAppSlotIdentityObservation struct {
	PrincipalID string `json:"principalId,omitempty" tf:"principal_id"`

	TenantID string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type FunctionAppSlotIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type FunctionAppSlotObservation struct {
	DefaultHostname string `json:"defaultHostname,omitempty" tf:"default_hostname"`

	Kind string `json:"kind,omitempty" tf:"kind"`

	OutboundIPAddresses string `json:"outboundIpAddresses,omitempty" tf:"outbound_ip_addresses"`

	PossibleOutboundIPAddresses string `json:"possibleOutboundIpAddresses,omitempty" tf:"possible_outbound_ip_addresses"`

	SiteCredential []FunctionAppSlotSiteCredentialObservation `json:"siteCredential,omitempty" tf:"site_credential"`
}

type FunctionAppSlotParameters struct {

	// +kubebuilder:validation:Required
	AppServicePlanID string `json:"appServicePlanId" tf:"app_service_plan_id"`

	// +kubebuilder:validation:Optional
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings"`

	// +kubebuilder:validation:Optional
	AuthSettings []FunctionAppSlotAuthSettingsParameters `json:"authSettings,omitempty" tf:"auth_settings"`

	// +kubebuilder:validation:Optional
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled"`

	// +kubebuilder:validation:Optional
	ConnectionString []FunctionAppSlotConnectionStringParameters `json:"connectionString,omitempty" tf:"connection_string"`

	// +kubebuilder:validation:Optional
	DailyMemoryTimeQuota *int64 `json:"dailyMemoryTimeQuota,omitempty" tf:"daily_memory_time_quota"`

	// +kubebuilder:validation:Optional
	EnableBuiltinLogging *bool `json:"enableBuiltinLogging,omitempty" tf:"enable_builtin_logging"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Required
	FunctionAppName string `json:"functionAppName" tf:"function_app_name"`

	// +kubebuilder:validation:Optional
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only"`

	// +kubebuilder:validation:Optional
	Identity []FunctionAppSlotIdentityParameters `json:"identity,omitempty" tf:"identity"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SiteConfig []FunctionAppSlotSiteConfigParameters `json:"siteConfig,omitempty" tf:"site_config"`

	// +kubebuilder:validation:Required
	StorageAccountAccessKey string `json:"storageAccountAccessKey" tf:"storage_account_access_key"`

	// +kubebuilder:validation:Required
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type FunctionAppSlotSiteConfigObservation struct {
}

type FunctionAppSlotSiteConfigParameters struct {

	// +kubebuilder:validation:Optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on"`

	// +kubebuilder:validation:Optional
	AppScaleLimit *int64 `json:"appScaleLimit,omitempty" tf:"app_scale_limit"`

	// +kubebuilder:validation:Optional
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty" tf:"auto_swap_slot_name"`

	// +kubebuilder:validation:Optional
	Cors []SiteConfigCorsParameters `json:"cors,omitempty" tf:"cors"`

	// +kubebuilder:validation:Optional
	DotnetFrameworkVersion *string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version"`

	// +kubebuilder:validation:Optional
	ElasticInstanceMinimum *int64 `json:"elasticInstanceMinimum,omitempty" tf:"elastic_instance_minimum"`

	// +kubebuilder:validation:Optional
	FtpsState *string `json:"ftpsState,omitempty" tf:"ftps_state"`

	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path"`

	// +kubebuilder:validation:Optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled"`

	// +kubebuilder:validation:Optional
	IPRestriction []SiteConfigIPRestrictionParameters `json:"ipRestriction,omitempty" tf:"ip_restriction"`

	// +kubebuilder:validation:Optional
	JavaVersion *string `json:"javaVersion,omitempty" tf:"java_version"`

	// +kubebuilder:validation:Optional
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version"`

	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`

	// +kubebuilder:validation:Optional
	PreWarmedInstanceCount *int64 `json:"preWarmedInstanceCount,omitempty" tf:"pre_warmed_instance_count"`

	// +kubebuilder:validation:Optional
	RuntimeScaleMonitoringEnabled *bool `json:"runtimeScaleMonitoringEnabled,omitempty" tf:"runtime_scale_monitoring_enabled"`

	// +kubebuilder:validation:Optional
	ScmIPRestriction []SiteConfigScmIPRestrictionParameters `json:"scmIpRestriction,omitempty" tf:"scm_ip_restriction"`

	// +kubebuilder:validation:Optional
	ScmType *string `json:"scmType,omitempty" tf:"scm_type"`

	// +kubebuilder:validation:Optional
	ScmUseMainIPRestriction *bool `json:"scmUseMainIpRestriction,omitempty" tf:"scm_use_main_ip_restriction"`

	// +kubebuilder:validation:Optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process"`

	// +kubebuilder:validation:Optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled"`
}

type FunctionAppSlotSiteCredentialObservation struct {
	Password string `json:"password,omitempty" tf:"password"`

	Username string `json:"username,omitempty" tf:"username"`
}

type FunctionAppSlotSiteCredentialParameters struct {
}

type IPRestrictionHeadersObservation struct {
}

type IPRestrictionHeadersParameters struct {

	// +kubebuilder:validation:Optional
	XAzureFdid []string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`

	// +kubebuilder:validation:Optional
	XFdHealthProbe []string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`

	// +kubebuilder:validation:Optional
	XForwardedFor []string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`

	// +kubebuilder:validation:Optional
	XForwardedHost []string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type SiteConfigCorsObservation struct {
}

type SiteConfigCorsParameters struct {

	// +kubebuilder:validation:Required
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`

	// +kubebuilder:validation:Optional
	SupportCredentials *bool `json:"supportCredentials,omitempty" tf:"support_credentials"`
}

type SiteConfigIPRestrictionObservation struct {
}

type SiteConfigIPRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// +kubebuilder:validation:Optional
	Headers []IPRestrictionHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	// +kubebuilder:validation:Optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id"`
}

type SiteConfigScmIPRestrictionHeadersObservation struct {
}

type SiteConfigScmIPRestrictionHeadersParameters struct {

	// +kubebuilder:validation:Optional
	XAzureFdid []string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`

	// +kubebuilder:validation:Optional
	XFdHealthProbe []string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`

	// +kubebuilder:validation:Optional
	XForwardedFor []string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`

	// +kubebuilder:validation:Optional
	XForwardedHost []string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type SiteConfigScmIPRestrictionObservation struct {
}

type SiteConfigScmIPRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// +kubebuilder:validation:Optional
	Headers []SiteConfigScmIPRestrictionHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	// +kubebuilder:validation:Optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id"`
}

// FunctionAppSlotSpec defines the desired state of FunctionAppSlot
type FunctionAppSlotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FunctionAppSlotParameters `json:"forProvider"`
}

// FunctionAppSlotStatus defines the observed state of FunctionAppSlot.
type FunctionAppSlotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FunctionAppSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppSlot is the Schema for the FunctionAppSlots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
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
	FunctionAppSlotKind             = "FunctionAppSlot"
	FunctionAppSlotGroupKind        = schema.GroupKind{Group: Group, Kind: FunctionAppSlotKind}.String()
	FunctionAppSlotKindAPIVersion   = FunctionAppSlotKind + "." + GroupVersion.String()
	FunctionAppSlotGroupVersionKind = GroupVersion.WithKind(FunctionAppSlotKind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppSlot{}, &FunctionAppSlotList{})
}
