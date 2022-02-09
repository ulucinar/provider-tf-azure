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

import (
	"context"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-azure/apis/azure2/v1alpha2"
	v1alpha21 "github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2"
	rconfig "github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this HBaseCluster.
func (mg *HBaseCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this HadoopCluster.
func (mg *HadoopCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this InteractiveQueryCluster.
func (mg *InteractiveQueryCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this KafkaCluster.
func (mg *KafkaCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].KafkaManagementNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].KafkaManagementNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].KafkaManagementNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].KafkaManagementNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].KafkaManagementNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].KafkaManagementNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].KafkaManagementNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this MLServicesCluster.
func (mg *MLServicesCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].EdgeNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this RServerCluster.
func (mg *RServerCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].EdgeNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].EdgeNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this SparkCluster.
func (mg *SparkCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this StormCluster.
func (mg *StormCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].HeadNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].HeadNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].WorkerNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].WorkerNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Roles); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Roles[i3].ZookeeperNode); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha21.SubnetList{},
					Managed: &v1alpha21.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID")
			}
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Roles[i3].ZookeeperNode[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}
