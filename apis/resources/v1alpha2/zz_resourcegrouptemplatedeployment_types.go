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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ResourceGroupTemplateDeploymentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OutputContent *string `json:"outputContent,omitempty" tf:"output_content,omitempty"`
}

type ResourceGroupTemplateDeploymentParameters struct {

	// +kubebuilder:validation:Optional
	DebugLevel *string `json:"debugLevel,omitempty" tf:"debug_level,omitempty"`

	// +kubebuilder:validation:Required
	DeploymentMode *string `json:"deploymentMode" tf:"deployment_mode,omitempty"`

	// +kubebuilder:validation:Optional
	ParametersContent *string `json:"parametersContent,omitempty" tf:"parameters_content,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateSpecVersionID *string `json:"templateSpecVersionId,omitempty" tf:"template_spec_version_id,omitempty"`
}

// ResourceGroupTemplateDeploymentSpec defines the desired state of ResourceGroupTemplateDeployment
type ResourceGroupTemplateDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupTemplateDeploymentParameters `json:"forProvider"`
}

// ResourceGroupTemplateDeploymentStatus defines the observed state of ResourceGroupTemplateDeployment.
type ResourceGroupTemplateDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupTemplateDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupTemplateDeployment is the Schema for the ResourceGroupTemplateDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ResourceGroupTemplateDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupTemplateDeploymentSpec   `json:"spec"`
	Status            ResourceGroupTemplateDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupTemplateDeploymentList contains a list of ResourceGroupTemplateDeployments
type ResourceGroupTemplateDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupTemplateDeployment `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupTemplateDeployment_Kind             = "ResourceGroupTemplateDeployment"
	ResourceGroupTemplateDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroupTemplateDeployment_Kind}.String()
	ResourceGroupTemplateDeployment_KindAPIVersion   = ResourceGroupTemplateDeployment_Kind + "." + CRDGroupVersion.String()
	ResourceGroupTemplateDeployment_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroupTemplateDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupTemplateDeployment{}, &ResourceGroupTemplateDeploymentList{})
}
