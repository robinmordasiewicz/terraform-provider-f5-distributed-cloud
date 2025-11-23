// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package forward_proxy_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ForwardProxyPolicyResourceModel struct {
	Name        types.String     `tfsdk:"name"`
	Namespace   types.String     `tfsdk:"namespace"`
	Description types.String     `tfsdk:"description"`
	ProxyRules  []ProxyRuleModel `tfsdk:"proxy_rules"`
	ID          types.String     `tfsdk:"id"`
}

type ProxyRuleModel struct {
	Action      types.String `tfsdk:"action"`
	DomainMatch types.String `tfsdk:"domain_match"`
	HTTPMethods types.List   `tfsdk:"http_methods"`
	PortMatch   types.Int64  `tfsdk:"port_match"`
}

type APIForwardProxyPolicy struct {
	Metadata   APIMetadata               `json:"metadata"`
	Spec       APIForwardProxyPolicySpec `json:"spec"`
	SystemMeta APISystemMetadata         `json:"system_metadata,omitempty"`
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

type APIForwardProxyPolicySpec struct {
	ProxyRules []APIProxyRule `json:"proxy_rules,omitempty"`
}

type APIProxyRule struct {
	Action      string   `json:"action,omitempty"`
	DomainMatch string   `json:"domain_match,omitempty"`
	HTTPMethods []string `json:"http_methods,omitempty"`
	PortMatch   int64    `json:"port_match,omitempty"`
}

func (m *ForwardProxyPolicyResourceModel) ToAPIRequest() *APIForwardProxyPolicy {
	spec := APIForwardProxyPolicySpec{}
	if len(m.ProxyRules) > 0 {
		spec.ProxyRules = make([]APIProxyRule, len(m.ProxyRules))
		for i, rule := range m.ProxyRules {
			apiRule := APIProxyRule{
				Action:      rule.Action.ValueString(),
				DomainMatch: rule.DomainMatch.ValueString(),
				PortMatch:   rule.PortMatch.ValueInt64(),
			}
			if !rule.HTTPMethods.IsNull() {
				methods := make([]string, 0)
				rule.HTTPMethods.ElementsAs(context.Background(), &methods, false)
				apiRule.HTTPMethods = methods
			}
			spec.ProxyRules[i] = apiRule
		}
	}
	return &APIForwardProxyPolicy{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *ForwardProxyPolicyResourceModel) FromAPIResponse(resp *APIForwardProxyPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
