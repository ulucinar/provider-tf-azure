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

type HighAvailabilityObservation struct {
}

type HighAvailabilityParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`

	// +kubebuilder:validation:Optional
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty" tf:"standby_availability_zone"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	DayOfWeek *int64 `json:"dayOfWeek,omitempty" tf:"day_of_week"`

	// +kubebuilder:validation:Optional
	StartHour *int64 `json:"startHour,omitempty" tf:"start_hour"`

	// +kubebuilder:validation:Optional
	StartMinute *int64 `json:"startMinute,omitempty" tf:"start_minute"`
}

type PostgresqlFlexibleServerObservation struct {
	CmkEnabled string `json:"cmkEnabled,omitempty" tf:"cmk_enabled"`

	Fqdn string `json:"fqdn,omitempty" tf:"fqdn"`

	PublicNetworkAccessEnabled bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
}

type PostgresqlFlexibleServerParameters struct {

	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login"`

	// +kubebuilder:validation:Optional
	AdministratorPassword *string `json:"administratorPassword,omitempty" tf:"administrator_password"`

	// +kubebuilder:validation:Optional
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`

	// +kubebuilder:validation:Optional
	DelegatedSubnetID *string `json:"delegatedSubnetId,omitempty" tf:"delegated_subnet_id"`

	// +kubebuilder:validation:Optional
	HighAvailability []HighAvailabilityParameters `json:"highAvailability,omitempty" tf:"high_availability"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc"`

	// +kubebuilder:validation:Optional
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name"`

	// +kubebuilder:validation:Optional
	SourceServerID *string `json:"sourceServerId,omitempty" tf:"source_server_id"`

	// +kubebuilder:validation:Optional
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

// PostgresqlFlexibleServerSpec defines the desired state of PostgresqlFlexibleServer
type PostgresqlFlexibleServerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostgresqlFlexibleServerParameters `json:"forProvider"`
}

// PostgresqlFlexibleServerStatus defines the observed state of PostgresqlFlexibleServer.
type PostgresqlFlexibleServerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostgresqlFlexibleServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlFlexibleServer is the Schema for the PostgresqlFlexibleServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlFlexibleServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlFlexibleServerSpec   `json:"spec"`
	Status            PostgresqlFlexibleServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlFlexibleServerList contains a list of PostgresqlFlexibleServers
type PostgresqlFlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlFlexibleServer `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlFlexibleServerKind             = "PostgresqlFlexibleServer"
	PostgresqlFlexibleServerGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlFlexibleServerKind}.String()
	PostgresqlFlexibleServerKindAPIVersion   = PostgresqlFlexibleServerKind + "." + GroupVersion.String()
	PostgresqlFlexibleServerGroupVersionKind = GroupVersion.WithKind(PostgresqlFlexibleServerKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlFlexibleServer{}, &PostgresqlFlexibleServerList{})
}
