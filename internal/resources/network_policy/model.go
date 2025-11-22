// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package network_policy

import "github.com/hashicorp/terraform-plugin-framework/types"

type NetworkPolicyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	PolicyType  types.String `tfsdk:"policy_type"`
	Rules       []NetworkRuleModel `tfsdk:"rules"`
	ID          types.String `tfsdk:"id"`
}

type NetworkRuleModel struct {
	Name         types.String `tfsdk:"name"`
	Action       types.String `tfsdk:"action"`
	SourcePrefix types.String `tfsdk:"source_prefix"`
	DestPrefix   types.String `tfsdk:"dest_prefix"`
	Protocol     types.String `tfsdk:"protocol"`
	Ports        types.String `tfsdk:"ports"`
}

type APINetworkPolicy struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APINetworkPolicySpec   `json:"spec"`
	SystemMeta APISystemMetadata      `json:"system_metadata,omitempty"`
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

type APINetworkPolicySpec struct {
	PolicyType string           `json:"policy_type,omitempty"`
	Rules      []APINetworkRule `json:"rules,omitempty"`
}

type APINetworkRule struct {
	Name         string `json:"name,omitempty"`
	Action       string `json:"action,omitempty"`
	SourcePrefix string `json:"source_prefix,omitempty"`
	DestPrefix   string `json:"dest_prefix,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	Ports        string `json:"ports,omitempty"`
}

func (m *NetworkPolicyResourceModel) ToAPIRequest() *APINetworkPolicy {
	spec := APINetworkPolicySpec{}
	if !m.PolicyType.IsNull() {
		spec.PolicyType = m.PolicyType.ValueString()
	}
	if len(m.Rules) > 0 {
		spec.Rules = make([]APINetworkRule, len(m.Rules))
		for i, rule := range m.Rules {
			spec.Rules[i] = APINetworkRule{
				Name:         rule.Name.ValueString(),
				Action:       rule.Action.ValueString(),
				SourcePrefix: rule.SourcePrefix.ValueString(),
				DestPrefix:   rule.DestPrefix.ValueString(),
				Protocol:     rule.Protocol.ValueString(),
				Ports:        rule.Ports.ValueString(),
			}
		}
	}
	return &APINetworkPolicy{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *NetworkPolicyResourceModel) FromAPIResponse(resp *APINetworkPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.PolicyType != "" {
		m.PolicyType = types.StringValue(resp.Spec.PolicyType)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
