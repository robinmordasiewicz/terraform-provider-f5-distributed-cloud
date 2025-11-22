// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package enhanced_firewall_policy

import "github.com/hashicorp/terraform-plugin-framework/types"

type EnhancedFirewallPolicyResourceModel struct {
	Name        types.String        `tfsdk:"name"`
	Namespace   types.String        `tfsdk:"namespace"`
	Description types.String        `tfsdk:"description"`
	Rules       []FirewallRuleModel `tfsdk:"rules"`
	ID          types.String        `tfsdk:"id"`
}

type FirewallRuleModel struct {
	Name     types.String `tfsdk:"name"`
	Action   types.String `tfsdk:"action"`
	Protocol types.String `tfsdk:"protocol"`
	SrcIP    types.String `tfsdk:"src_ip"`
	DstIP    types.String `tfsdk:"dst_ip"`
	SrcPort  types.Int64  `tfsdk:"src_port"`
	DstPort  types.Int64  `tfsdk:"dst_port"`
}

type APIEnhancedFirewallPolicy struct {
	Metadata   APIMetadata                   `json:"metadata"`
	Spec       APIEnhancedFirewallPolicySpec `json:"spec"`
	SystemMeta APISystemMetadata             `json:"system_metadata,omitempty"`
}

type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
}

type APIEnhancedFirewallPolicySpec struct {
	Rules []APIFirewallRule `json:"rules,omitempty"`
}

type APIFirewallRule struct {
	Name     string `json:"name,omitempty"`
	Action   string `json:"action,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	SrcIP    string `json:"src_ip,omitempty"`
	DstIP    string `json:"dst_ip,omitempty"`
	SrcPort  int64  `json:"src_port,omitempty"`
	DstPort  int64  `json:"dst_port,omitempty"`
}

func (m *EnhancedFirewallPolicyResourceModel) ToAPIRequest() *APIEnhancedFirewallPolicy {
	spec := APIEnhancedFirewallPolicySpec{}
	if len(m.Rules) > 0 {
		spec.Rules = make([]APIFirewallRule, len(m.Rules))
		for i, rule := range m.Rules {
			spec.Rules[i] = APIFirewallRule{
				Name:     rule.Name.ValueString(),
				Action:   rule.Action.ValueString(),
				Protocol: rule.Protocol.ValueString(),
				SrcIP:    rule.SrcIP.ValueString(),
				DstIP:    rule.DstIP.ValueString(),
				SrcPort:  rule.SrcPort.ValueInt64(),
				DstPort:  rule.DstPort.ValueInt64(),
			}
		}
	}
	return &APIEnhancedFirewallPolicy{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *EnhancedFirewallPolicyResourceModel) FromAPIResponse(resp *APIEnhancedFirewallPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
