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

type ContainerGroupObservation struct {
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn"`

	IPAddress string `json:"ipAddress,omitempty" tf:"ip_address"`
}

type ContainerGroupParameters struct {

	// +kubebuilder:validation:Required
	Container []ContainerParameters `json:"container" tf:"container"`

	// +kubebuilder:validation:Optional
	DNSConfig []DNSConfigParameters `json:"dnsConfig,omitempty" tf:"dns_config"`

	// +kubebuilder:validation:Optional
	DNSNameLabel *string `json:"dnsNameLabel,omitempty" tf:"dns_name_label"`

	// +kubebuilder:validation:Optional
	Diagnostics []DiagnosticsParameters `json:"diagnostics,omitempty" tf:"diagnostics"`

	// +kubebuilder:validation:Optional
	ExposedPort []ExposedPortParameters `json:"exposedPort,omitempty" tf:"exposed_port"`

	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	// +kubebuilder:validation:Optional
	ImageRegistryCredential []ImageRegistryCredentialParameters `json:"imageRegistryCredential,omitempty" tf:"image_registry_credential"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NetworkProfileID *string `json:"networkProfileId,omitempty" tf:"network_profile_id"`

	// +kubebuilder:validation:Required
	OsType string `json:"osType" tf:"os_type"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	RestartPolicy *string `json:"restartPolicy,omitempty" tf:"restart_policy"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ContainerObservation struct {
}

type ContainerParameters struct {

	// +kubebuilder:validation:Required
	CPU float64 `json:"cpu" tf:"cpu"`

	// +kubebuilder:validation:Optional
	Commands []string `json:"commands,omitempty" tf:"commands"`

	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`

	// +kubebuilder:validation:Optional
	Gpu []GpuParameters `json:"gpu,omitempty" tf:"gpu"`

	// +kubebuilder:validation:Required
	Image string `json:"image" tf:"image"`

	// +kubebuilder:validation:Optional
	LivenessProbe []LivenessProbeParameters `json:"livenessProbe,omitempty" tf:"liveness_probe"`

	// +kubebuilder:validation:Required
	Memory float64 `json:"memory" tf:"memory"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Ports []PortsParameters `json:"ports,omitempty" tf:"ports"`

	// +kubebuilder:validation:Optional
	ReadinessProbe []ReadinessProbeParameters `json:"readinessProbe,omitempty" tf:"readiness_probe"`

	// +kubebuilder:validation:Optional
	SecureEnvironmentVariables map[string]string `json:"secureEnvironmentVariables,omitempty" tf:"secure_environment_variables"`

	// +kubebuilder:validation:Optional
	Volume []VolumeParameters `json:"volume,omitempty" tf:"volume"`
}

type DNSConfigObservation struct {
}

type DNSConfigParameters struct {

	// +kubebuilder:validation:Required
	Nameservers []string `json:"nameservers" tf:"nameservers"`

	// +kubebuilder:validation:Required
	Options []string `json:"options" tf:"options"`

	// +kubebuilder:validation:Required
	SearchDomains []string `json:"searchDomains" tf:"search_domains"`
}

type DiagnosticsObservation struct {
}

type DiagnosticsParameters struct {

	// +kubebuilder:validation:Required
	LogAnalytics []LogAnalyticsParameters `json:"logAnalytics" tf:"log_analytics"`
}

type ExposedPortObservation struct {
}

type ExposedPortParameters struct {

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
}

type GitRepoObservation struct {
}

type GitRepoParameters struct {

	// +kubebuilder:validation:Optional
	Directory *string `json:"directory,omitempty" tf:"directory"`

	// +kubebuilder:validation:Optional
	Revision *string `json:"revision,omitempty" tf:"revision"`

	// +kubebuilder:validation:Required
	URL string `json:"url" tf:"url"`
}

type GpuObservation struct {
}

type GpuParameters struct {

	// +kubebuilder:validation:Optional
	Count *int64 `json:"count,omitempty" tf:"count"`

	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku"`
}

type HTTPGetObservation struct {
}

type HTTPGetParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type IdentityObservation struct {
	PrincipalID string `json:"principalId,omitempty" tf:"principal_id"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type ImageRegistryCredentialObservation struct {
}

type ImageRegistryCredentialParameters struct {

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Server string `json:"server" tf:"server"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

type LivenessProbeObservation struct {
}

type LivenessProbeParameters struct {

	// +kubebuilder:validation:Optional
	Exec []string `json:"exec,omitempty" tf:"exec"`

	// +kubebuilder:validation:Optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`

	// +kubebuilder:validation:Optional
	HTTPGet []HTTPGetParameters `json:"httpGet,omitempty" tf:"http_get"`

	// +kubebuilder:validation:Optional
	InitialDelaySeconds *int64 `json:"initialDelaySeconds,omitempty" tf:"initial_delay_seconds"`

	// +kubebuilder:validation:Optional
	PeriodSeconds *int64 `json:"periodSeconds,omitempty" tf:"period_seconds"`

	// +kubebuilder:validation:Optional
	SuccessThreshold *int64 `json:"successThreshold,omitempty" tf:"success_threshold"`

	// +kubebuilder:validation:Optional
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds"`
}

type LogAnalyticsObservation struct {
}

type LogAnalyticsParameters struct {

	// +kubebuilder:validation:Optional
	LogType *string `json:"logType,omitempty" tf:"log_type"`

	// +kubebuilder:validation:Optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata"`

	// +kubebuilder:validation:Required
	WorkspaceID string `json:"workspaceId" tf:"workspace_id"`

	// +kubebuilder:validation:Required
	WorkspaceKey string `json:"workspaceKey" tf:"workspace_key"`
}

type PortsObservation struct {
}

type PortsParameters struct {

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
}

type ReadinessProbeHTTPGetObservation struct {
}

type ReadinessProbeHTTPGetParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type ReadinessProbeObservation struct {
}

type ReadinessProbeParameters struct {

	// +kubebuilder:validation:Optional
	Exec []string `json:"exec,omitempty" tf:"exec"`

	// +kubebuilder:validation:Optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`

	// +kubebuilder:validation:Optional
	HTTPGet []ReadinessProbeHTTPGetParameters `json:"httpGet,omitempty" tf:"http_get"`

	// +kubebuilder:validation:Optional
	InitialDelaySeconds *int64 `json:"initialDelaySeconds,omitempty" tf:"initial_delay_seconds"`

	// +kubebuilder:validation:Optional
	PeriodSeconds *int64 `json:"periodSeconds,omitempty" tf:"period_seconds"`

	// +kubebuilder:validation:Optional
	SuccessThreshold *int64 `json:"successThreshold,omitempty" tf:"success_threshold"`

	// +kubebuilder:validation:Optional
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds"`
}

type VolumeObservation struct {
}

type VolumeParameters struct {

	// +kubebuilder:validation:Optional
	EmptyDir *bool `json:"emptyDir,omitempty" tf:"empty_dir"`

	// +kubebuilder:validation:Optional
	GitRepo []GitRepoParameters `json:"gitRepo,omitempty" tf:"git_repo"`

	// +kubebuilder:validation:Required
	MountPath string `json:"mountPath" tf:"mount_path"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`

	// +kubebuilder:validation:Optional
	Secret map[string]string `json:"secret,omitempty" tf:"secret"`

	// +kubebuilder:validation:Optional
	ShareName *string `json:"shareName,omitempty" tf:"share_name"`

	// +kubebuilder:validation:Optional
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key"`

	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name"`
}

// ContainerGroupSpec defines the desired state of ContainerGroup
type ContainerGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ContainerGroupParameters `json:"forProvider"`
}

// ContainerGroupStatus defines the observed state of ContainerGroup.
type ContainerGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ContainerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerGroup is the Schema for the ContainerGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ContainerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerGroupSpec   `json:"spec"`
	Status            ContainerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerGroupList contains a list of ContainerGroups
type ContainerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerGroup `json:"items"`
}

// Repository type metadata.
var (
	ContainerGroupKind             = "ContainerGroup"
	ContainerGroupGroupKind        = schema.GroupKind{Group: Group, Kind: ContainerGroupKind}.String()
	ContainerGroupKindAPIVersion   = ContainerGroupKind + "." + GroupVersion.String()
	ContainerGroupGroupVersionKind = GroupVersion.WithKind(ContainerGroupKind)
)

func init() {
	SchemeBuilder.Register(&ContainerGroup{}, &ContainerGroupList{})
}
