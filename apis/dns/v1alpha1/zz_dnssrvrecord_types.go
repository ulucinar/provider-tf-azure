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

type DnsSrvRecordObservation struct {
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn"`
}

type DnsSrvRecordParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Record []DnsSrvRecordRecordParameters `json:"record" tf:"record"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	TTL int64 `json:"ttl" tf:"ttl"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	ZoneName string `json:"zoneName" tf:"zone_name"`
}

type DnsSrvRecordRecordObservation struct {
}

type DnsSrvRecordRecordParameters struct {

	// +kubebuilder:validation:Required
	Port int64 `json:"port" tf:"port"`

	// +kubebuilder:validation:Required
	Priority int64 `json:"priority" tf:"priority"`

	// +kubebuilder:validation:Required
	Target string `json:"target" tf:"target"`

	// +kubebuilder:validation:Required
	Weight int64 `json:"weight" tf:"weight"`
}

// DnsSrvRecordSpec defines the desired state of DnsSrvRecord
type DnsSrvRecordSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DnsSrvRecordParameters `json:"forProvider"`
}

// DnsSrvRecordStatus defines the observed state of DnsSrvRecord.
type DnsSrvRecordStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DnsSrvRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DnsSrvRecord is the Schema for the DnsSrvRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DnsSrvRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsSrvRecordSpec   `json:"spec"`
	Status            DnsSrvRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnsSrvRecordList contains a list of DnsSrvRecords
type DnsSrvRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsSrvRecord `json:"items"`
}

// Repository type metadata.
var (
	DnsSrvRecordKind             = "DnsSrvRecord"
	DnsSrvRecordGroupKind        = schema.GroupKind{Group: Group, Kind: DnsSrvRecordKind}.String()
	DnsSrvRecordKindAPIVersion   = DnsSrvRecordKind + "." + GroupVersion.String()
	DnsSrvRecordGroupVersionKind = GroupVersion.WithKind(DnsSrvRecordKind)
)

func init() {
	SchemeBuilder.Register(&DnsSrvRecord{}, &DnsSrvRecordList{})
}
