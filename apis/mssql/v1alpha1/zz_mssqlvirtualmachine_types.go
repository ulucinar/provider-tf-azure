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

type AutoBackupObservation struct {
}

type AutoBackupParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty" tf:"encryption_enabled"`

	// +kubebuilder:validation:Optional
	EncryptionPassword *string `json:"encryptionPassword,omitempty" tf:"encryption_password"`

	// +kubebuilder:validation:Optional
	ManualSchedule []ManualScheduleParameters `json:"manualSchedule,omitempty" tf:"manual_schedule"`

	// +kubebuilder:validation:Required
	RetentionPeriodInDays int64 `json:"retentionPeriodInDays" tf:"retention_period_in_days"`

	// +kubebuilder:validation:Required
	StorageAccountAccessKey string `json:"storageAccountAccessKey" tf:"storage_account_access_key"`

	// +kubebuilder:validation:Required
	StorageBlobEndpoint string `json:"storageBlobEndpoint" tf:"storage_blob_endpoint"`

	// +kubebuilder:validation:Optional
	SystemDatabasesBackupEnabled *bool `json:"systemDatabasesBackupEnabled,omitempty" tf:"system_databases_backup_enabled"`
}

type AutoPatchingObservation struct {
}

type AutoPatchingParameters struct {

	// +kubebuilder:validation:Required
	DayOfWeek string `json:"dayOfWeek" tf:"day_of_week"`

	// +kubebuilder:validation:Required
	MaintenanceWindowDurationInMinutes int64 `json:"maintenanceWindowDurationInMinutes" tf:"maintenance_window_duration_in_minutes"`

	// +kubebuilder:validation:Required
	MaintenanceWindowStartingHour int64 `json:"maintenanceWindowStartingHour" tf:"maintenance_window_starting_hour"`
}

type DataSettingsObservation struct {
}

type DataSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultFilePath string `json:"defaultFilePath" tf:"default_file_path"`

	// +kubebuilder:validation:Required
	Luns []int64 `json:"luns" tf:"luns"`
}

type KeyVaultCredentialObservation struct {
}

type KeyVaultCredentialParameters struct {

	// +kubebuilder:validation:Required
	KeyVaultURL string `json:"keyVaultUrl" tf:"key_vault_url"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ServicePrincipalName string `json:"servicePrincipalName" tf:"service_principal_name"`

	// +kubebuilder:validation:Required
	ServicePrincipalSecret string `json:"servicePrincipalSecret" tf:"service_principal_secret"`
}

type LogSettingsObservation struct {
}

type LogSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultFilePath string `json:"defaultFilePath" tf:"default_file_path"`

	// +kubebuilder:validation:Required
	Luns []int64 `json:"luns" tf:"luns"`
}

type ManualScheduleObservation struct {
}

type ManualScheduleParameters struct {

	// +kubebuilder:validation:Required
	FullBackupFrequency string `json:"fullBackupFrequency" tf:"full_backup_frequency"`

	// +kubebuilder:validation:Required
	FullBackupStartHour int64 `json:"fullBackupStartHour" tf:"full_backup_start_hour"`

	// +kubebuilder:validation:Required
	FullBackupWindowInHours int64 `json:"fullBackupWindowInHours" tf:"full_backup_window_in_hours"`

	// +kubebuilder:validation:Required
	LogBackupFrequencyInMinutes int64 `json:"logBackupFrequencyInMinutes" tf:"log_backup_frequency_in_minutes"`
}

type MssqlVirtualMachineObservation struct {
}

type MssqlVirtualMachineParameters struct {

	// +kubebuilder:validation:Optional
	AutoBackup []AutoBackupParameters `json:"autoBackup,omitempty" tf:"auto_backup"`

	// +kubebuilder:validation:Optional
	AutoPatching []AutoPatchingParameters `json:"autoPatching,omitempty" tf:"auto_patching"`

	// +kubebuilder:validation:Optional
	KeyVaultCredential []KeyVaultCredentialParameters `json:"keyVaultCredential,omitempty" tf:"key_vault_credential"`

	// +kubebuilder:validation:Optional
	RServicesEnabled *bool `json:"rServicesEnabled,omitempty" tf:"r_services_enabled"`

	// +kubebuilder:validation:Optional
	SQLConnectivityPort *int64 `json:"sqlConnectivityPort,omitempty" tf:"sql_connectivity_port"`

	// +kubebuilder:validation:Optional
	SQLConnectivityType *string `json:"sqlConnectivityType,omitempty" tf:"sql_connectivity_type"`

	// +kubebuilder:validation:Optional
	SQLConnectivityUpdatePassword *string `json:"sqlConnectivityUpdatePassword,omitempty" tf:"sql_connectivity_update_password"`

	// +kubebuilder:validation:Optional
	SQLConnectivityUpdateUsername *string `json:"sqlConnectivityUpdateUsername,omitempty" tf:"sql_connectivity_update_username"`

	// +kubebuilder:validation:Required
	SQLLicenseType string `json:"sqlLicenseType" tf:"sql_license_type"`

	// +kubebuilder:validation:Optional
	StorageConfiguration []StorageConfigurationParameters `json:"storageConfiguration,omitempty" tf:"storage_configuration"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	VirtualMachineID string `json:"virtualMachineId" tf:"virtual_machine_id"`
}

type StorageConfigurationObservation struct {
}

type StorageConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	DataSettings []DataSettingsParameters `json:"dataSettings,omitempty" tf:"data_settings"`

	// +kubebuilder:validation:Required
	DiskType string `json:"diskType" tf:"disk_type"`

	// +kubebuilder:validation:Optional
	LogSettings []LogSettingsParameters `json:"logSettings,omitempty" tf:"log_settings"`

	// +kubebuilder:validation:Required
	StorageWorkloadType string `json:"storageWorkloadType" tf:"storage_workload_type"`

	// +kubebuilder:validation:Optional
	TempDBSettings []TempDBSettingsParameters `json:"tempDbSettings,omitempty" tf:"temp_db_settings"`
}

type TempDBSettingsObservation struct {
}

type TempDBSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultFilePath string `json:"defaultFilePath" tf:"default_file_path"`

	// +kubebuilder:validation:Required
	Luns []int64 `json:"luns" tf:"luns"`
}

// MssqlVirtualMachineSpec defines the desired state of MssqlVirtualMachine
type MssqlVirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MssqlVirtualMachineParameters `json:"forProvider"`
}

// MssqlVirtualMachineStatus defines the observed state of MssqlVirtualMachine.
type MssqlVirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MssqlVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MssqlVirtualMachine is the Schema for the MssqlVirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MssqlVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlVirtualMachineSpec   `json:"spec"`
	Status            MssqlVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MssqlVirtualMachineList contains a list of MssqlVirtualMachines
type MssqlVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MssqlVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	MssqlVirtualMachineKind             = "MssqlVirtualMachine"
	MssqlVirtualMachineGroupKind        = schema.GroupKind{Group: Group, Kind: MssqlVirtualMachineKind}.String()
	MssqlVirtualMachineKindAPIVersion   = MssqlVirtualMachineKind + "." + GroupVersion.String()
	MssqlVirtualMachineGroupVersionKind = GroupVersion.WithKind(MssqlVirtualMachineKind)
)

func init() {
	SchemeBuilder.Register(&MssqlVirtualMachine{}, &MssqlVirtualMachineList{})
}
