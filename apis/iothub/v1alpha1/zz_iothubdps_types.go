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

type IothubDpsObservation struct {
	DeviceProvisioningHostName string `json:"deviceProvisioningHostName,omitempty" tf:"device_provisioning_host_name"`

	IDScope string `json:"idScope,omitempty" tf:"id_scope"`

	ServiceOperationsHostName string `json:"serviceOperationsHostName,omitempty" tf:"service_operations_host_name"`
}

type IothubDpsParameters struct {

	// +kubebuilder:validation:Optional
	AllocationPolicy *string `json:"allocationPolicy,omitempty" tf:"allocation_policy"`

	// +kubebuilder:validation:Optional
	LinkedHub []LinkedHubParameters `json:"linkedHub,omitempty" tf:"linked_hub"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	Sku []IothubDpsSkuParameters `json:"sku" tf:"sku"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type IothubDpsSkuObservation struct {
}

type IothubDpsSkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity int64 `json:"capacity" tf:"capacity"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type LinkedHubObservation struct {
	Hostname string `json:"hostname,omitempty" tf:"hostname"`
}

type LinkedHubParameters struct {

	// +kubebuilder:validation:Optional
	AllocationWeight *int64 `json:"allocationWeight,omitempty" tf:"allocation_weight"`

	// +kubebuilder:validation:Optional
	ApplyAllocationPolicy *bool `json:"applyAllocationPolicy,omitempty" tf:"apply_allocation_policy"`

	// +kubebuilder:validation:Required
	ConnectionString string `json:"connectionString" tf:"connection_string"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`
}

// IothubDpsSpec defines the desired state of IothubDps
type IothubDpsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IothubDpsParameters `json:"forProvider"`
}

// IothubDpsStatus defines the observed state of IothubDps.
type IothubDpsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IothubDpsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IothubDps is the Schema for the IothubDpss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IothubDps struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubDpsSpec   `json:"spec"`
	Status            IothubDpsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubDpsList contains a list of IothubDpss
type IothubDpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IothubDps `json:"items"`
}

// Repository type metadata.
var (
	IothubDpsKind             = "IothubDps"
	IothubDpsGroupKind        = schema.GroupKind{Group: Group, Kind: IothubDpsKind}.String()
	IothubDpsKindAPIVersion   = IothubDpsKind + "." + GroupVersion.String()
	IothubDpsGroupVersionKind = GroupVersion.WithKind(IothubDpsKind)
)

func init() {
	SchemeBuilder.Register(&IothubDps{}, &IothubDpsList{})
}
