// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Association) DeepCopyInto(out *Association) {
	*out = *in
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(SubnetIDReferencerForRouteTable)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Association.
func (in *Association) DeepCopy() *Association {
	if in == nil {
		return nil
	}
	out := new(Association)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationState) DeepCopyInto(out *AssociationState) {
	*out = *in
	in.Association.DeepCopyInto(&out.Association)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationState.
func (in *AssociationState) DeepCopy() *AssociationState {
	if in == nil {
		return nil
	}
	out := new(AssociationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPermission) DeepCopyInto(out *IPPermission) {
	*out = *in
	if in.CIDRBlocks != nil {
		in, out := &in.CIDRBlocks, &out.CIDRBlocks
		*out = make([]IPRange, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPermission.
func (in *IPPermission) DeepCopy() *IPPermission {
	if in == nil {
		return nil
	}
	out := new(IPPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRange) DeepCopyInto(out *IPRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRange.
func (in *IPRange) DeepCopy() *IPRange {
	if in == nil {
		return nil
	}
	out := new(IPRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGateway) DeepCopyInto(out *InternetGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGateway.
func (in *InternetGateway) DeepCopy() *InternetGateway {
	if in == nil {
		return nil
	}
	out := new(InternetGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternetGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayAttachment) DeepCopyInto(out *InternetGatewayAttachment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayAttachment.
func (in *InternetGatewayAttachment) DeepCopy() *InternetGatewayAttachment {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayExternalStatus) DeepCopyInto(out *InternetGatewayExternalStatus) {
	*out = *in
	if in.Attachments != nil {
		in, out := &in.Attachments, &out.Attachments
		*out = make([]InternetGatewayAttachment, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayExternalStatus.
func (in *InternetGatewayExternalStatus) DeepCopy() *InternetGatewayExternalStatus {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayIDReferencer) DeepCopyInto(out *InternetGatewayIDReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayIDReferencer.
func (in *InternetGatewayIDReferencer) DeepCopy() *InternetGatewayIDReferencer {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayIDReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayIDReferencerForRouteTable) DeepCopyInto(out *InternetGatewayIDReferencerForRouteTable) {
	*out = *in
	out.InternetGatewayIDReferencer = in.InternetGatewayIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayIDReferencerForRouteTable.
func (in *InternetGatewayIDReferencerForRouteTable) DeepCopy() *InternetGatewayIDReferencerForRouteTable {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayIDReferencerForRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayList) DeepCopyInto(out *InternetGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InternetGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayList.
func (in *InternetGatewayList) DeepCopy() *InternetGatewayList {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternetGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayParameters) DeepCopyInto(out *InternetGatewayParameters) {
	*out = *in
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(VPCIDReferencerForInternetGateway)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayParameters.
func (in *InternetGatewayParameters) DeepCopy() *InternetGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewaySpec) DeepCopyInto(out *InternetGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.InternetGatewayParameters.DeepCopyInto(&out.InternetGatewayParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewaySpec.
func (in *InternetGatewaySpec) DeepCopy() *InternetGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(InternetGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternetGatewayStatus) DeepCopyInto(out *InternetGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.InternetGatewayExternalStatus.DeepCopyInto(&out.InternetGatewayExternalStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternetGatewayStatus.
func (in *InternetGatewayStatus) DeepCopy() *InternetGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(InternetGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.GatewayIDRef != nil {
		in, out := &in.GatewayIDRef, &out.GatewayIDRef
		*out = new(InternetGatewayIDReferencerForRouteTable)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteState) DeepCopyInto(out *RouteState) {
	*out = *in
	in.Route.DeepCopyInto(&out.Route)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteState.
func (in *RouteState) DeepCopy() *RouteState {
	if in == nil {
		return nil
	}
	out := new(RouteState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTable.
func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableExternalStatus) DeepCopyInto(out *RouteTableExternalStatus) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Associations != nil {
		in, out := &in.Associations, &out.Associations
		*out = make([]AssociationState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableExternalStatus.
func (in *RouteTableExternalStatus) DeepCopy() *RouteTableExternalStatus {
	if in == nil {
		return nil
	}
	out := new(RouteTableExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableList) DeepCopyInto(out *RouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableList.
func (in *RouteTableList) DeepCopy() *RouteTableList {
	if in == nil {
		return nil
	}
	out := new(RouteTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableParameters) DeepCopyInto(out *RouteTableParameters) {
	*out = *in
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(VPCIDReferencerForRouteTable)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Associations != nil {
		in, out := &in.Associations, &out.Associations
		*out = make([]Association, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableParameters.
func (in *RouteTableParameters) DeepCopy() *RouteTableParameters {
	if in == nil {
		return nil
	}
	out := new(RouteTableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableSpec) DeepCopyInto(out *RouteTableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.RouteTableParameters.DeepCopyInto(&out.RouteTableParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableSpec.
func (in *RouteTableSpec) DeepCopy() *RouteTableSpec {
	if in == nil {
		return nil
	}
	out := new(RouteTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableStatus) DeepCopyInto(out *RouteTableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.RouteTableExternalStatus.DeepCopyInto(&out.RouteTableExternalStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableStatus.
func (in *RouteTableStatus) DeepCopy() *RouteTableStatus {
	if in == nil {
		return nil
	}
	out := new(RouteTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupExternalStatus) DeepCopyInto(out *SecurityGroupExternalStatus) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupExternalStatus.
func (in *SecurityGroupExternalStatus) DeepCopy() *SecurityGroupExternalStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupIDReferencer) DeepCopyInto(out *SecurityGroupIDReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupIDReferencer.
func (in *SecurityGroupIDReferencer) DeepCopy() *SecurityGroupIDReferencer {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupIDReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupList) DeepCopyInto(out *SecurityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupList.
func (in *SecurityGroupList) DeepCopy() *SecurityGroupList {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupParameters) DeepCopyInto(out *SecurityGroupParameters) {
	*out = *in
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(VPCIDReferencerForSecurityGroup)
		**out = **in
	}
	if in.IngressPermissions != nil {
		in, out := &in.IngressPermissions, &out.IngressPermissions
		*out = make([]IPPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EgressPermissions != nil {
		in, out := &in.EgressPermissions, &out.EgressPermissions
		*out = make([]IPPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupParameters.
func (in *SecurityGroupParameters) DeepCopy() *SecurityGroupParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupSpec) DeepCopyInto(out *SecurityGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.SecurityGroupParameters.DeepCopyInto(&out.SecurityGroupParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupSpec.
func (in *SecurityGroupSpec) DeepCopy() *SecurityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupStatus) DeepCopyInto(out *SecurityGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.SecurityGroupExternalStatus.DeepCopyInto(&out.SecurityGroupExternalStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupStatus.
func (in *SecurityGroupStatus) DeepCopy() *SecurityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetExternalStatus) DeepCopyInto(out *SubnetExternalStatus) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetExternalStatus.
func (in *SubnetExternalStatus) DeepCopy() *SubnetExternalStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetIDReferencer) DeepCopyInto(out *SubnetIDReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetIDReferencer.
func (in *SubnetIDReferencer) DeepCopy() *SubnetIDReferencer {
	if in == nil {
		return nil
	}
	out := new(SubnetIDReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetIDReferencerForRouteTable) DeepCopyInto(out *SubnetIDReferencerForRouteTable) {
	*out = *in
	out.SubnetIDReferencer = in.SubnetIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetIDReferencerForRouteTable.
func (in *SubnetIDReferencerForRouteTable) DeepCopy() *SubnetIDReferencerForRouteTable {
	if in == nil {
		return nil
	}
	out := new(SubnetIDReferencerForRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetList) DeepCopyInto(out *SubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetList.
func (in *SubnetList) DeepCopy() *SubnetList {
	if in == nil {
		return nil
	}
	out := new(SubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetParameters) DeepCopyInto(out *SubnetParameters) {
	*out = *in
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(VPCIDReferencerForSubnet)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetParameters.
func (in *SubnetParameters) DeepCopy() *SubnetParameters {
	if in == nil {
		return nil
	}
	out := new(SubnetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.SubnetParameters.DeepCopyInto(&out.SubnetParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.SubnetExternalStatus.DeepCopyInto(&out.SubnetExternalStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCExternalStatus) DeepCopyInto(out *VPCExternalStatus) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCExternalStatus.
func (in *VPCExternalStatus) DeepCopy() *VPCExternalStatus {
	if in == nil {
		return nil
	}
	out := new(VPCExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCIDReferencer) DeepCopyInto(out *VPCIDReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCIDReferencer.
func (in *VPCIDReferencer) DeepCopy() *VPCIDReferencer {
	if in == nil {
		return nil
	}
	out := new(VPCIDReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCIDReferencerForInternetGateway) DeepCopyInto(out *VPCIDReferencerForInternetGateway) {
	*out = *in
	out.VPCIDReferencer = in.VPCIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCIDReferencerForInternetGateway.
func (in *VPCIDReferencerForInternetGateway) DeepCopy() *VPCIDReferencerForInternetGateway {
	if in == nil {
		return nil
	}
	out := new(VPCIDReferencerForInternetGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCIDReferencerForRouteTable) DeepCopyInto(out *VPCIDReferencerForRouteTable) {
	*out = *in
	out.VPCIDReferencer = in.VPCIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCIDReferencerForRouteTable.
func (in *VPCIDReferencerForRouteTable) DeepCopy() *VPCIDReferencerForRouteTable {
	if in == nil {
		return nil
	}
	out := new(VPCIDReferencerForRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCIDReferencerForSecurityGroup) DeepCopyInto(out *VPCIDReferencerForSecurityGroup) {
	*out = *in
	out.VPCIDReferencer = in.VPCIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCIDReferencerForSecurityGroup.
func (in *VPCIDReferencerForSecurityGroup) DeepCopy() *VPCIDReferencerForSecurityGroup {
	if in == nil {
		return nil
	}
	out := new(VPCIDReferencerForSecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCIDReferencerForSubnet) DeepCopyInto(out *VPCIDReferencerForSubnet) {
	*out = *in
	out.VPCIDReferencer = in.VPCIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCIDReferencerForSubnet.
func (in *VPCIDReferencerForSubnet) DeepCopy() *VPCIDReferencerForSubnet {
	if in == nil {
		return nil
	}
	out := new(VPCIDReferencerForSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCList) DeepCopyInto(out *VPCList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCList.
func (in *VPCList) DeepCopy() *VPCList {
	if in == nil {
		return nil
	}
	out := new(VPCList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCParameters) DeepCopyInto(out *VPCParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCParameters.
func (in *VPCParameters) DeepCopy() *VPCParameters {
	if in == nil {
		return nil
	}
	out := new(VPCParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSpec) DeepCopyInto(out *VPCSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.VPCParameters = in.VPCParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSpec.
func (in *VPCSpec) DeepCopy() *VPCSpec {
	if in == nil {
		return nil
	}
	out := new(VPCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCStatus) DeepCopyInto(out *VPCStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.VPCExternalStatus.DeepCopyInto(&out.VPCExternalStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCStatus.
func (in *VPCStatus) DeepCopy() *VPCStatus {
	if in == nil {
		return nil
	}
	out := new(VPCStatus)
	in.DeepCopyInto(out)
	return out
}