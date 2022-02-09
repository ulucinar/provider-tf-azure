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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type JobScheduleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type JobScheduleParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	JobScheduleID *string `json:"jobScheduleId,omitempty" tf:"job_schedule_id,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	RunBookName *string `json:"runbookName" tf:"runbook_name,omitempty"`

	// +kubebuilder:validation:Optional
	RunOn *string `json:"runOn,omitempty" tf:"run_on,omitempty"`

	// +kubebuilder:validation:Required
	ScheduleName *string `json:"scheduleName" tf:"schedule_name,omitempty"`
}

// JobScheduleSpec defines the desired state of JobSchedule
type JobScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobScheduleParameters `json:"forProvider"`
}

// JobScheduleStatus defines the observed state of JobSchedule.
type JobScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobSchedule is the Schema for the JobSchedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type JobSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobScheduleSpec   `json:"spec"`
	Status            JobScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobScheduleList contains a list of JobSchedules
type JobScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobSchedule `json:"items"`
}

// Repository type metadata.
var (
	JobSchedule_Kind             = "JobSchedule"
	JobSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobSchedule_Kind}.String()
	JobSchedule_KindAPIVersion   = JobSchedule_Kind + "." + CRDGroupVersion.String()
	JobSchedule_GroupVersionKind = CRDGroupVersion.WithKind(JobSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&JobSchedule{}, &JobScheduleList{})
}
