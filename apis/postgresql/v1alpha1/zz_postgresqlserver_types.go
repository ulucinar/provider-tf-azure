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

type IdentityObservation struct {
	PrincipalID string `json:"principalId,omitempty" tf:"principal_id"`

	TenantID string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type PostgresqlServerObservation struct {
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn"`
}

type PostgresqlServerParameters struct {

	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login"`

	// +kubebuilder:validation:Optional
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty" tf:"administrator_login_password"`

	// +kubebuilder:validation:Optional
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled"`

	// +kubebuilder:validation:Optional
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`

	// +kubebuilder:validation:Optional
	CreationSourceServerID *string `json:"creationSourceServerId,omitempty" tf:"creation_source_server_id"`

	// +kubebuilder:validation:Optional
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	// +kubebuilder:validation:Optional
	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time"`

	// +kubebuilder:validation:Required
	SkuName string `json:"skuName" tf:"sku_name"`

	// +kubebuilder:validation:Optional
	SslEnforcement *string `json:"sslEnforcement,omitempty" tf:"ssl_enforcement"`

	// +kubebuilder:validation:Optional
	SslEnforcementEnabled *bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled"`

	// +kubebuilder:validation:Optional
	SslMinimalTLSVersionEnforced *string `json:"sslMinimalTlsVersionEnforced,omitempty" tf:"ssl_minimal_tls_version_enforced"`

	// +kubebuilder:validation:Optional
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`

	// +kubebuilder:validation:Optional
	StorageProfile []StorageProfileParameters `json:"storageProfile,omitempty" tf:"storage_profile"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	ThreatDetectionPolicy []ThreatDetectionPolicyParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy"`

	// +kubebuilder:validation:Required
	Version string `json:"version" tf:"version"`
}

type StorageProfileObservation struct {
}

type StorageProfileParameters struct {

	// +kubebuilder:validation:Optional
	AutoGrow *string `json:"autoGrow,omitempty" tf:"auto_grow"`

	// +kubebuilder:validation:Optional
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`

	// +kubebuilder:validation:Optional
	GeoRedundantBackup *string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup"`

	// +kubebuilder:validation:Optional
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`
}

type ThreatDetectionPolicyObservation struct {
}

type ThreatDetectionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DisabledAlerts []string `json:"disabledAlerts,omitempty" tf:"disabled_alerts"`

	// +kubebuilder:validation:Optional
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty" tf:"email_account_admins"`

	// +kubebuilder:validation:Optional
	EmailAddresses []string `json:"emailAddresses,omitempty" tf:"email_addresses"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKey *string `json:"storageAccountAccessKey,omitempty" tf:"storage_account_access_key"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint"`
}

// PostgresqlServerSpec defines the desired state of PostgresqlServer
type PostgresqlServerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostgresqlServerParameters `json:"forProvider"`
}

// PostgresqlServerStatus defines the observed state of PostgresqlServer.
type PostgresqlServerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostgresqlServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServer is the Schema for the PostgresqlServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlServerSpec   `json:"spec"`
	Status            PostgresqlServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlServerList contains a list of PostgresqlServers
type PostgresqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlServer `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlServerKind             = "PostgresqlServer"
	PostgresqlServerGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlServerKind}.String()
	PostgresqlServerKindAPIVersion   = PostgresqlServerKind + "." + GroupVersion.String()
	PostgresqlServerGroupVersionKind = GroupVersion.WithKind(PostgresqlServerKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlServer{}, &PostgresqlServerList{})
}
