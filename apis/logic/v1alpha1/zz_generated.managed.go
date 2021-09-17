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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppActionCustom.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppActionCustom) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppActionCustom.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppActionCustom) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppActionCustom.
func (mg *LogicAppActionCustom) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppActionHttp.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppActionHttp) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppActionHttp.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppActionHttp) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppActionHttp.
func (mg *LogicAppActionHttp) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppIntegrationAccount.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccount) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppIntegrationAccount.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccount) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppIntegrationAccount.
func (mg *LogicAppIntegrationAccount) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppIntegrationAccountCertificate.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccountCertificate) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppIntegrationAccountCertificate.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccountCertificate) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppIntegrationAccountCertificate.
func (mg *LogicAppIntegrationAccountCertificate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppIntegrationAccountSchema.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccountSchema) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppIntegrationAccountSchema.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccountSchema) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppIntegrationAccountSchema.
func (mg *LogicAppIntegrationAccountSchema) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppIntegrationAccountSession.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccountSession) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppIntegrationAccountSession.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppIntegrationAccountSession) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppIntegrationAccountSession.
func (mg *LogicAppIntegrationAccountSession) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppTriggerCustom.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppTriggerCustom) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppTriggerCustom.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppTriggerCustom) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppTriggerCustom.
func (mg *LogicAppTriggerCustom) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppTriggerHttpRequest.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppTriggerHttpRequest) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppTriggerHttpRequest.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppTriggerHttpRequest) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppTriggerHttpRequest.
func (mg *LogicAppTriggerHttpRequest) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppTriggerRecurrence.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppTriggerRecurrence) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppTriggerRecurrence.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppTriggerRecurrence) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppTriggerRecurrence.
func (mg *LogicAppTriggerRecurrence) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicAppWorkflow.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicAppWorkflow) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicAppWorkflow.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicAppWorkflow) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this LogicAppWorkflow.
func (mg *LogicAppWorkflow) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
