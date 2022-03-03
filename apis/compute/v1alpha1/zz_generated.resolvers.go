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
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CubeServer.
func (mg *CubeServer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.DatacenterCfg.DatacenterID,
		Extract:      ExtractDatacenterID(),
		Reference:    mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef,
		Selector:     mg.Spec.ForProvider.DatacenterCfg.DatacenterIDSelector,
		To: reference.To{
			List:    &DatacenterList{},
			Managed: &Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatacenterCfg.DatacenterID")
	}
	mg.Spec.ForProvider.DatacenterCfg.DatacenterID = rsp.ResolvedValue
	mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Lan.
func (mg *Lan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.DatacenterCfg.DatacenterID,
		Extract:      ExtractDatacenterID(),
		Reference:    mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef,
		Selector:     mg.Spec.ForProvider.DatacenterCfg.DatacenterIDSelector,
		To: reference.To{
			List:    &DatacenterList{},
			Managed: &Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatacenterCfg.DatacenterID")
	}
	mg.Spec.ForProvider.DatacenterCfg.DatacenterID = rsp.ResolvedValue
	mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.DatacenterCfg.DatacenterID,
		Extract:      ExtractDatacenterID(),
		Reference:    mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef,
		Selector:     mg.Spec.ForProvider.DatacenterCfg.DatacenterIDSelector,
		To: reference.To{
			List:    &DatacenterList{},
			Managed: &Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatacenterCfg.DatacenterID")
	}
	mg.Spec.ForProvider.DatacenterCfg.DatacenterID = rsp.ResolvedValue
	mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.VolumeCfg.VolumeID,
		Extract:      ExtractVolumeID(),
		Reference:    mg.Spec.ForProvider.VolumeCfg.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeCfg.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeCfg.VolumeID")
	}
	mg.Spec.ForProvider.VolumeCfg.VolumeID = rsp.ResolvedValue
	mg.Spec.ForProvider.VolumeCfg.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Volume.
func (mg *Volume) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.DatacenterCfg.DatacenterID,
		Extract:      ExtractDatacenterID(),
		Reference:    mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef,
		Selector:     mg.Spec.ForProvider.DatacenterCfg.DatacenterIDSelector,
		To: reference.To{
			List:    &DatacenterList{},
			Managed: &Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatacenterCfg.DatacenterID")
	}
	mg.Spec.ForProvider.DatacenterCfg.DatacenterID = rsp.ResolvedValue
	mg.Spec.ForProvider.DatacenterCfg.DatacenterIDRef = rsp.ResolvedReference

	return nil
}
