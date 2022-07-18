/*
Copyright 2022 The Crossplane Authors.

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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this IPGroup
func (mg *IPGroup) GetTerraformResourceType() string {
	return "azurerm_ip_group"
}

// GetConnectionDetailsMapping for this IPGroup
func (tr *IPGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this IPGroup
func (tr *IPGroup) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this IPGroup
func (tr *IPGroup) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this IPGroup
func (tr *IPGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this IPGroup
func (tr *IPGroup) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this IPGroup
func (tr *IPGroup) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this IPGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *IPGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &IPGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *IPGroup) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LoadBalancer
func (mg *LoadBalancer) GetTerraformResourceType() string {
	return "azurerm_lb"
}

// GetConnectionDetailsMapping for this LoadBalancer
func (tr *LoadBalancer) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LoadBalancer
func (tr *LoadBalancer) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LoadBalancer
func (tr *LoadBalancer) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LoadBalancer
func (tr *LoadBalancer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LoadBalancer
func (tr *LoadBalancer) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LoadBalancer
func (tr *LoadBalancer) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LoadBalancer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LoadBalancer) LateInitialize(attrs []byte) (bool, error) {
	params := &LoadBalancerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LoadBalancer) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NetworkInterface
func (mg *NetworkInterface) GetTerraformResourceType() string {
	return "azurerm_network_interface"
}

// GetConnectionDetailsMapping for this NetworkInterface
func (tr *NetworkInterface) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NetworkInterface
func (tr *NetworkInterface) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NetworkInterface
func (tr *NetworkInterface) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NetworkInterface
func (tr *NetworkInterface) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NetworkInterface
func (tr *NetworkInterface) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NetworkInterface
func (tr *NetworkInterface) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NetworkInterface using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NetworkInterface) LateInitialize(attrs []byte) (bool, error) {
	params := &NetworkInterfaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NetworkInterface) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this PublicIP
func (mg *PublicIP) GetTerraformResourceType() string {
	return "azurerm_public_ip"
}

// GetConnectionDetailsMapping for this PublicIP
func (tr *PublicIP) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this PublicIP
func (tr *PublicIP) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PublicIP
func (tr *PublicIP) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this PublicIP
func (tr *PublicIP) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this PublicIP
func (tr *PublicIP) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PublicIP
func (tr *PublicIP) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this PublicIP using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PublicIP) LateInitialize(attrs []byte) (bool, error) {
	params := &PublicIPParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *PublicIP) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Subnet
func (mg *Subnet) GetTerraformResourceType() string {
	return "azurerm_subnet"
}

// GetConnectionDetailsMapping for this Subnet
func (tr *Subnet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Subnet
func (tr *Subnet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Subnet
func (tr *Subnet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Subnet
func (tr *Subnet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Subnet
func (tr *Subnet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Subnet
func (tr *Subnet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Subnet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Subnet) LateInitialize(attrs []byte) (bool, error) {
	params := &SubnetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("AddressPrefix"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Subnet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubnetNATGatewayAssociation
func (mg *SubnetNATGatewayAssociation) GetTerraformResourceType() string {
	return "azurerm_subnet_nat_gateway_association"
}

// GetConnectionDetailsMapping for this SubnetNATGatewayAssociation
func (tr *SubnetNATGatewayAssociation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubnetNATGatewayAssociation
func (tr *SubnetNATGatewayAssociation) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubnetNATGatewayAssociation
func (tr *SubnetNATGatewayAssociation) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubnetNATGatewayAssociation
func (tr *SubnetNATGatewayAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubnetNATGatewayAssociation
func (tr *SubnetNATGatewayAssociation) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubnetNATGatewayAssociation
func (tr *SubnetNATGatewayAssociation) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubnetNATGatewayAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubnetNATGatewayAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &SubnetNATGatewayAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubnetNATGatewayAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubnetNetworkSecurityGroupAssociation
func (mg *SubnetNetworkSecurityGroupAssociation) GetTerraformResourceType() string {
	return "azurerm_subnet_network_security_group_association"
}

// GetConnectionDetailsMapping for this SubnetNetworkSecurityGroupAssociation
func (tr *SubnetNetworkSecurityGroupAssociation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubnetNetworkSecurityGroupAssociation
func (tr *SubnetNetworkSecurityGroupAssociation) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubnetNetworkSecurityGroupAssociation
func (tr *SubnetNetworkSecurityGroupAssociation) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubnetNetworkSecurityGroupAssociation
func (tr *SubnetNetworkSecurityGroupAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubnetNetworkSecurityGroupAssociation
func (tr *SubnetNetworkSecurityGroupAssociation) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubnetNetworkSecurityGroupAssociation
func (tr *SubnetNetworkSecurityGroupAssociation) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubnetNetworkSecurityGroupAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubnetNetworkSecurityGroupAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &SubnetNetworkSecurityGroupAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubnetNetworkSecurityGroupAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubnetRouteTableAssociation
func (mg *SubnetRouteTableAssociation) GetTerraformResourceType() string {
	return "azurerm_subnet_route_table_association"
}

// GetConnectionDetailsMapping for this SubnetRouteTableAssociation
func (tr *SubnetRouteTableAssociation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubnetRouteTableAssociation
func (tr *SubnetRouteTableAssociation) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubnetRouteTableAssociation
func (tr *SubnetRouteTableAssociation) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubnetRouteTableAssociation
func (tr *SubnetRouteTableAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubnetRouteTableAssociation
func (tr *SubnetRouteTableAssociation) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubnetRouteTableAssociation
func (tr *SubnetRouteTableAssociation) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubnetRouteTableAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubnetRouteTableAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &SubnetRouteTableAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubnetRouteTableAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubnetServiceEndpointStoragePolicy
func (mg *SubnetServiceEndpointStoragePolicy) GetTerraformResourceType() string {
	return "azurerm_subnet_service_endpoint_storage_policy"
}

// GetConnectionDetailsMapping for this SubnetServiceEndpointStoragePolicy
func (tr *SubnetServiceEndpointStoragePolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubnetServiceEndpointStoragePolicy
func (tr *SubnetServiceEndpointStoragePolicy) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubnetServiceEndpointStoragePolicy
func (tr *SubnetServiceEndpointStoragePolicy) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubnetServiceEndpointStoragePolicy
func (tr *SubnetServiceEndpointStoragePolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubnetServiceEndpointStoragePolicy
func (tr *SubnetServiceEndpointStoragePolicy) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubnetServiceEndpointStoragePolicy
func (tr *SubnetServiceEndpointStoragePolicy) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubnetServiceEndpointStoragePolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubnetServiceEndpointStoragePolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &SubnetServiceEndpointStoragePolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubnetServiceEndpointStoragePolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualNetwork
func (mg *VirtualNetwork) GetTerraformResourceType() string {
	return "azurerm_virtual_network"
}

// GetConnectionDetailsMapping for this VirtualNetwork
func (tr *VirtualNetwork) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VirtualNetwork
func (tr *VirtualNetwork) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetwork
func (tr *VirtualNetwork) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetwork
func (tr *VirtualNetwork) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetwork
func (tr *VirtualNetwork) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetwork
func (tr *VirtualNetwork) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetwork using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetwork) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("Subnet"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetwork) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualNetworkGateway
func (mg *VirtualNetworkGateway) GetTerraformResourceType() string {
	return "azurerm_virtual_network_gateway"
}

// GetConnectionDetailsMapping for this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetworkGateway using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetworkGateway) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkGatewayParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetworkGateway) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualNetworkGatewayConnection
func (mg *VirtualNetworkGatewayConnection) GetTerraformResourceType() string {
	return "azurerm_virtual_network_gateway_connection"
}

// GetConnectionDetailsMapping for this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"authorization_key": "spec.forProvider.authorizationKeySecretRef", "shared_key": "spec.forProvider.sharedKeySecretRef"}
}

// GetObservation of this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetworkGatewayConnection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetworkGatewayConnection) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkGatewayConnectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetworkGatewayConnection) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualNetworkPeering
func (mg *VirtualNetworkPeering) GetTerraformResourceType() string {
	return "azurerm_virtual_network_peering"
}

// GetConnectionDetailsMapping for this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetworkPeering using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetworkPeering) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkPeeringParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetworkPeering) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualWAN
func (mg *VirtualWAN) GetTerraformResourceType() string {
	return "azurerm_virtual_wan"
}

// GetConnectionDetailsMapping for this VirtualWAN
func (tr *VirtualWAN) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VirtualWAN
func (tr *VirtualWAN) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualWAN
func (tr *VirtualWAN) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualWAN
func (tr *VirtualWAN) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualWAN
func (tr *VirtualWAN) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualWAN
func (tr *VirtualWAN) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualWAN using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualWAN) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualWANParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualWAN) GetTerraformSchemaVersion() int {
	return 0
}
