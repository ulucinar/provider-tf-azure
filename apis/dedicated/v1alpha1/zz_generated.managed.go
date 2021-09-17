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

// GetCondition of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DedicatedHardwareSecurityModule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DedicatedHardwareSecurityModule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DedicatedHardwareSecurityModule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DedicatedHardwareSecurityModule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DedicatedHardwareSecurityModule.
func (mg *DedicatedHardwareSecurityModule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DedicatedHost.
func (mg *DedicatedHost) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DedicatedHost.
func (mg *DedicatedHost) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DedicatedHost.
func (mg *DedicatedHost) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DedicatedHost.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DedicatedHost) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DedicatedHost.
func (mg *DedicatedHost) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DedicatedHost.
func (mg *DedicatedHost) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DedicatedHost.
func (mg *DedicatedHost) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DedicatedHost.
func (mg *DedicatedHost) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DedicatedHost.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DedicatedHost) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DedicatedHost.
func (mg *DedicatedHost) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
