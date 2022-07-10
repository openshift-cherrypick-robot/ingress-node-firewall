//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewall) DeepCopyInto(out *IngressNodeFirewall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewall.
func (in *IngressNodeFirewall) DeepCopy() *IngressNodeFirewall {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressNodeFirewall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallFromCIDR) DeepCopyInto(out *IngressNodeFirewallFromCIDR) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallFromCIDR.
func (in *IngressNodeFirewallFromCIDR) DeepCopy() *IngressNodeFirewallFromCIDR {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallFromCIDR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallICMPRule) DeepCopyInto(out *IngressNodeFirewallICMPRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallICMPRule.
func (in *IngressNodeFirewallICMPRule) DeepCopy() *IngressNodeFirewallICMPRule {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallICMPRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallList) DeepCopyInto(out *IngressNodeFirewallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressNodeFirewall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallList.
func (in *IngressNodeFirewallList) DeepCopy() *IngressNodeFirewallList {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressNodeFirewallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallProtoRule) DeepCopyInto(out *IngressNodeFirewallProtoRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallProtoRule.
func (in *IngressNodeFirewallProtoRule) DeepCopy() *IngressNodeFirewallProtoRule {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallProtoRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallProtocolRule) DeepCopyInto(out *IngressNodeFirewallProtocolRule) {
	*out = *in
	in.ProtocolRule.DeepCopyInto(&out.ProtocolRule)
	out.ICMPRule = in.ICMPRule
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallProtocolRule.
func (in *IngressNodeFirewallProtocolRule) DeepCopy() *IngressNodeFirewallProtocolRule {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallProtocolRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallRules) DeepCopyInto(out *IngressNodeFirewallRules) {
	*out = *in
	if in.FromCIDRs != nil {
		in, out := &in.FromCIDRs, &out.FromCIDRs
		*out = make([]IngressNodeFirewallFromCIDR, len(*in))
		copy(*out, *in)
	}
	if in.FirewallProtocolRules != nil {
		in, out := &in.FirewallProtocolRules, &out.FirewallProtocolRules
		*out = make([]IngressNodeFirewallProtocolRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallRules.
func (in *IngressNodeFirewallRules) DeepCopy() *IngressNodeFirewallRules {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallSpec) DeepCopyInto(out *IngressNodeFirewallSpec) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]IngressNodeFirewallRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallSpec.
func (in *IngressNodeFirewallSpec) DeepCopy() *IngressNodeFirewallSpec {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNodeFirewallStatus) DeepCopyInto(out *IngressNodeFirewallStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNodeFirewallStatus.
func (in *IngressNodeFirewallStatus) DeepCopy() *IngressNodeFirewallStatus {
	if in == nil {
		return nil
	}
	out := new(IngressNodeFirewallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeEndpoint) DeepCopyInto(out *NodeEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeEndpoint.
func (in *NodeEndpoint) DeepCopy() *NodeEndpoint {
	if in == nil {
		return nil
	}
	out := new(NodeEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeEndpointList) DeepCopyInto(out *NodeEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeEndpointList.
func (in *NodeEndpointList) DeepCopy() *NodeEndpointList {
	if in == nil {
		return nil
	}
	out := new(NodeEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeEndpointSpec) DeepCopyInto(out *NodeEndpointSpec) {
	*out = *in
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeEndpointSpec.
func (in *NodeEndpointSpec) DeepCopy() *NodeEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(NodeEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeEndpointStatus) DeepCopyInto(out *NodeEndpointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeEndpointStatus.
func (in *NodeEndpointStatus) DeepCopy() *NodeEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(NodeEndpointStatus)
	in.DeepCopyInto(out)
	return out
}