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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IngressNodeFirewallFromCIDR define ingress node CIDR matching field
type IngressNodeFirewallFromCIDR struct {
	// FromCIDR is the IPv4/IPv6 CIDR from which we apply node firewall rule
	FromCIDR string `json:"fromCIDR" protobuf:"bytes,1,opt,name=fromCIDR"`
}

// IngressNodeFirewallICMPRule define ingress node firewall rule for ICMP and ICMPv6 protocols
type IngressNodeFirewallICMPRule struct {
	// ICMPType define ICMP Type Numbers (RFC 792).
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Maximum:=43
	// +kubebuilder:validation:Minimum:=0
	ICMPType int8 `json:"icmpType" protobuf:"varint,1,opt,name=icmpType"`

	// ICMPCode define ICMP Code ID (RFC 792).
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Maximum:=16
	// +kubebuilder:validation:Minimum:=0
	ICMPCode int8 `json:"icmpCode" protobuf:"varint,2,opt,name=icmpCode"`
}

// IngressNodeFirewallProtoRule define ingress node firewall rule for TCP, UDP and SCTP protocols
type IngressNodeFirewallProtoRule struct {
	// start-end range of dstPorts or [port1, port2,..., portn].
	// +kubebuilder:validation:Optional
	// +optional
	Ports []int32 `json:"ports" protobuf:"varint,1,opt,name=ports"`
}

// IngressNodeFirewallProtocolRule define ingress node firewall rule per protocol
type IngressNodeFirewallProtocolRule struct {
	// IngressNodeFirewallProtoRule define ingress node firewall rule for TCP, UDP and SCTP protocols.
	ProtocolRule IngressNodeFirewallProtoRule `json:"protoRule" protobuf:"varint,1,opt,name=protoRule"`

	// IngressNodeFirewallICMPRule define ingress node firewall rule for ICMP and ICMPv6 protocols.
	ICMPRule IngressNodeFirewallICMPRule `json:"icmpRule" protobuf:"varint,2,opt,name:icmpRule"`

	// Protocol can be ICMP, ICMPv6, TCP, SCTP or UDP.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum="ICMP";"ICMPv6";"TCP";"UDP";"SCTP"
	Protocol IngressNodeFirewallRuleProtocolType `json:"protocol" protobuf:"varint,3,opt,name:protocol"`

	// Action can be Allow or Deny.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum="Allow";"Deny"
	Action IngressNodeFirewallActionType `json:"action" protobuf:"bytes,4,opt,name=action,casttype=IngressNodeFirewallActionType"`
}

// ProtocolType defines the protocol types that are supported
type IngressNodeFirewallRuleProtocolType string

const (
	// ProtocolTypeICMP refers to the ICMP protocol
	ProtocolTypeICMP IngressNodeFirewallRuleProtocolType = "ICMP"

	// ProtocolTypeICMPv6 refers to the ICMPv6 protocol
	ProtocolTypeICMPv6 IngressNodeFirewallRuleProtocolType = "ICMPv6"

	// ProtocolTypeTCP refers to the TCP protocol, for either IPv4 or IPv6
	ProtocolTypeTCP IngressNodeFirewallRuleProtocolType = "TCP"

	// ProtocolTypeUDP refers to the UDP protocol, for either IPv4 or IPv6
	ProtocolTypeUDP IngressNodeFirewallRuleProtocolType = "UDP"

	// ProtocolTypeSCTP refers to the SCTP protocol, for either IPv4 or IPv6
	ProtocolTypeSCTP IngressNodeFirewallRuleProtocolType = "SCTP"
)

// IngressNodeFirewallActionType indicates whether an IngressNodeFirewallRule allows or denies traffic
// +kubebuilder:validation:Pattern=`^Allow|Deny$`
type IngressNodeFirewallActionType string

const (
	IngressNodeFirewallAllow IngressNodeFirewallActionType = "Allow"
	DIngressNodeFirewallDeny IngressNodeFirewallActionType = "Deny"
)

// IngressNodeFirewallRules define ingress node firewall rule
type IngressNodeFirewallRules struct {
	// FromCIDRS is A list of CIDR from which we apply node firewall rule
	FromCIDRs []IngressNodeFirewallFromCIDR `json:"fromCIDRs" protobuf:"bytes,1,opt,name=fromCIDRs"`
	// FirewallProtocolRules is A list of per protocol ingress node firewall rules
	FirewallProtocolRules []IngressNodeFirewallProtocolRule `json:"firewallProtocolRules" protobuf:"bytes,2,opt,names=firewallProtocolRules"`
}

// IngressNodeFirewallSpec defines the desired state of IngressNodeFirewall
type IngressNodeFirewallSpec struct {
	// A list of ingress firewall policy rules.
	// empty list indicates no ingress firewall i.e allow all incoming traffic.
	// +kubebuilder:validation:Optional
	// +optional
	Ingress []IngressNodeFirewallRules `json:"ingress" protobuf:"bytes,2,opt,name=ingress"`
}

// IngressNodeFirewallStatus defines the observed state of IngressNodeFirewall
type IngressNodeFirewallStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// IngressNodeFirewall is the Schema for the ingressnodefirewalls API
type IngressNodeFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IngressNodeFirewallSpec   `json:"spec,omitempty"`
	Status IngressNodeFirewallStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IngressNodeFirewallList contains a list of IngressNodeFirewall
type IngressNodeFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressNodeFirewall `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IngressNodeFirewall{}, &IngressNodeFirewallList{})
}