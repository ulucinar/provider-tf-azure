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

type DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordObservation struct {
}

type DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName string `json:"linkedServiceName" tf:"linked_service_name"`

	// +kubebuilder:validation:Required
	SecretName string `json:"secretName" tf:"secret_name"`
}

type DataFactoryLinkedServiceAzureDatabricksObservation struct {
}

type DataFactoryLinkedServiceAzureDatabricksParameters struct {

	// +kubebuilder:validation:Optional
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token"`

	// +kubebuilder:validation:Required
	AdbDomain string `json:"adbDomain" tf:"adb_domain"`

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`

	// +kubebuilder:validation:Optional
	Annotations []string `json:"annotations,omitempty" tf:"annotations"`

	// +kubebuilder:validation:Required
	DataFactoryName string `json:"dataFactoryName" tf:"data_factory_name"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	ExistingClusterID *string `json:"existingClusterId,omitempty" tf:"existing_cluster_id"`

	// +kubebuilder:validation:Optional
	InstancePool []InstancePoolParameters `json:"instancePool,omitempty" tf:"instance_pool"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name"`

	// +kubebuilder:validation:Optional
	KeyVaultPassword []DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password"`

	// +kubebuilder:validation:Optional
	MsiWorkSpaceResourceID *string `json:"msiWorkSpaceResourceId,omitempty" tf:"msi_work_space_resource_id"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NewClusterConfig []NewClusterConfigParameters `json:"newClusterConfig,omitempty" tf:"new_cluster_config"`

	// +kubebuilder:validation:Optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type InstancePoolObservation struct {
}

type InstancePoolParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	// +kubebuilder:validation:Required
	InstancePoolID string `json:"instancePoolId" tf:"instance_pool_id"`

	// +kubebuilder:validation:Optional
	MaxNumberOfWorkers *int64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers"`

	// +kubebuilder:validation:Optional
	MinNumberOfWorkers *int64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers"`
}

type NewClusterConfigObservation struct {
}

type NewClusterConfigParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`

	// +kubebuilder:validation:Optional
	CustomTags map[string]string `json:"customTags,omitempty" tf:"custom_tags"`

	// +kubebuilder:validation:Optional
	DriverNodeType *string `json:"driverNodeType,omitempty" tf:"driver_node_type"`

	// +kubebuilder:validation:Optional
	InitScripts []string `json:"initScripts,omitempty" tf:"init_scripts"`

	// +kubebuilder:validation:Optional
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination"`

	// +kubebuilder:validation:Optional
	MaxNumberOfWorkers *int64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers"`

	// +kubebuilder:validation:Optional
	MinNumberOfWorkers *int64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers"`

	// +kubebuilder:validation:Required
	NodeType string `json:"nodeType" tf:"node_type"`

	// +kubebuilder:validation:Optional
	SparkConfig map[string]string `json:"sparkConfig,omitempty" tf:"spark_config"`

	// +kubebuilder:validation:Optional
	SparkEnvironmentVariables map[string]string `json:"sparkEnvironmentVariables,omitempty" tf:"spark_environment_variables"`
}

// DataFactoryLinkedServiceAzureDatabricksSpec defines the desired state of DataFactoryLinkedServiceAzureDatabricks
type DataFactoryLinkedServiceAzureDatabricksSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataFactoryLinkedServiceAzureDatabricksParameters `json:"forProvider"`
}

// DataFactoryLinkedServiceAzureDatabricksStatus defines the observed state of DataFactoryLinkedServiceAzureDatabricks.
type DataFactoryLinkedServiceAzureDatabricksStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataFactoryLinkedServiceAzureDatabricksObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryLinkedServiceAzureDatabricks is the Schema for the DataFactoryLinkedServiceAzureDatabrickss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryLinkedServiceAzureDatabricks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServiceAzureDatabricksSpec   `json:"spec"`
	Status            DataFactoryLinkedServiceAzureDatabricksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryLinkedServiceAzureDatabricksList contains a list of DataFactoryLinkedServiceAzureDatabrickss
type DataFactoryLinkedServiceAzureDatabricksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryLinkedServiceAzureDatabricks `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryLinkedServiceAzureDatabricksKind             = "DataFactoryLinkedServiceAzureDatabricks"
	DataFactoryLinkedServiceAzureDatabricksGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryLinkedServiceAzureDatabricksKind}.String()
	DataFactoryLinkedServiceAzureDatabricksKindAPIVersion   = DataFactoryLinkedServiceAzureDatabricksKind + "." + GroupVersion.String()
	DataFactoryLinkedServiceAzureDatabricksGroupVersionKind = GroupVersion.WithKind(DataFactoryLinkedServiceAzureDatabricksKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryLinkedServiceAzureDatabricks{}, &DataFactoryLinkedServiceAzureDatabricksList{})
}
