// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package rbac_policy

import "github.com/hashicorp/terraform-plugin-framework/types"

type RBACPolicyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Rules       []RuleModel  `tfsdk:"rules"`
	ID          types.String `tfsdk:"id"`
}

type RuleModel struct {
	Action    types.String   `tfsdk:"action"`
	APIGroups []types.String `tfsdk:"api_groups"`
	Resources []types.String `tfsdk:"resources"`
	Verbs     []types.String `tfsdk:"verbs"`
}

type APIRBACPolicy struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIRBACPolicySpec `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
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

type APIRBACPolicySpec struct {
	Rules []APIRule `json:"rules,omitempty"`
}

type APIRule struct {
	Action    string   `json:"action,omitempty"`
	APIGroups []string `json:"api_groups,omitempty"`
	Resources []string `json:"resources,omitempty"`
	Verbs     []string `json:"verbs,omitempty"`
}

func (m *RBACPolicyResourceModel) ToAPIRequest() *APIRBACPolicy {
	spec := APIRBACPolicySpec{}
	if len(m.Rules) > 0 {
		rules := make([]APIRule, len(m.Rules))
		for i, r := range m.Rules {
			rule := APIRule{
				Action: r.Action.ValueString(),
			}
			if len(r.APIGroups) > 0 {
				apiGroups := make([]string, len(r.APIGroups))
				for j, g := range r.APIGroups {
					apiGroups[j] = g.ValueString()
				}
				rule.APIGroups = apiGroups
			}
			if len(r.Resources) > 0 {
				resources := make([]string, len(r.Resources))
				for j, res := range r.Resources {
					resources[j] = res.ValueString()
				}
				rule.Resources = resources
			}
			if len(r.Verbs) > 0 {
				verbs := make([]string, len(r.Verbs))
				for j, v := range r.Verbs {
					verbs[j] = v.ValueString()
				}
				rule.Verbs = verbs
			}
			rules[i] = rule
		}
		spec.Rules = rules
	}
	return &APIRBACPolicy{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *RBACPolicyResourceModel) FromAPIResponse(resp *APIRBACPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
