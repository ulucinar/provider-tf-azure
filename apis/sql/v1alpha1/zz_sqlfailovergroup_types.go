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

type PartnerServersObservation struct {
	Location string `json:"location,omitempty" tf:"location"`

	Role string `json:"role,omitempty" tf:"role"`
}

type PartnerServersParameters struct {

	// +kubebuilder:validation:Required
	ID string `json:"id" tf:"id"`
}

type ReadWriteEndpointFailoverPolicyObservation struct {
}

type ReadWriteEndpointFailoverPolicyParameters struct {

	// +kubebuilder:validation:Optional
	GraceMinutes *int64 `json:"graceMinutes,omitempty" tf:"grace_minutes"`

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`
}

type ReadonlyEndpointFailoverPolicyObservation struct {
}

type ReadonlyEndpointFailoverPolicyParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`
}

type SqlFailoverGroupObservation struct {
	Location string `json:"location,omitempty" tf:"location"`

	Role string `json:"role,omitempty" tf:"role"`
}

type SqlFailoverGroupParameters struct {

	// +kubebuilder:validation:Optional
	Databases []string `json:"databases,omitempty" tf:"databases"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	PartnerServers []PartnerServersParameters `json:"partnerServers" tf:"partner_servers"`

	// +kubebuilder:validation:Required
	ReadWriteEndpointFailoverPolicy []ReadWriteEndpointFailoverPolicyParameters `json:"readWriteEndpointFailoverPolicy" tf:"read_write_endpoint_failover_policy"`

	// +kubebuilder:validation:Optional
	ReadonlyEndpointFailoverPolicy []ReadonlyEndpointFailoverPolicyParameters `json:"readonlyEndpointFailoverPolicy,omitempty" tf:"readonly_endpoint_failover_policy"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Required
	ServerName string `json:"serverName" tf:"server_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// SqlFailoverGroupSpec defines the desired state of SqlFailoverGroup
type SqlFailoverGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SqlFailoverGroupParameters `json:"forProvider"`
}

// SqlFailoverGroupStatus defines the observed state of SqlFailoverGroup.
type SqlFailoverGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SqlFailoverGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFailoverGroup is the Schema for the SqlFailoverGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SqlFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlFailoverGroupSpec   `json:"spec"`
	Status            SqlFailoverGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFailoverGroupList contains a list of SqlFailoverGroups
type SqlFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlFailoverGroup `json:"items"`
}

// Repository type metadata.
var (
	SqlFailoverGroupKind             = "SqlFailoverGroup"
	SqlFailoverGroupGroupKind        = schema.GroupKind{Group: Group, Kind: SqlFailoverGroupKind}.String()
	SqlFailoverGroupKindAPIVersion   = SqlFailoverGroupKind + "." + GroupVersion.String()
	SqlFailoverGroupGroupVersionKind = GroupVersion.WithKind(SqlFailoverGroupKind)
)

func init() {
	SchemeBuilder.Register(&SqlFailoverGroup{}, &SqlFailoverGroupList{})
}
