// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package waf_rule

import "github.com/hashicorp/terraform-plugin-framework/types"

type WAFRuleResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Mode        types.String `tfsdk:"mode"`
	RuleID      types.String `tfsdk:"rule_id"`
	Action      types.String `tfsdk:"action"`
	ID          types.String `tfsdk:"id"`
}

type APIWAFRule struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIWAFRuleSpec    `json:"spec"`
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

type APIWAFRuleSpec struct {
	Mode   string `json:"mode,omitempty"`
	RuleID string `json:"rule_id,omitempty"`
	Action string `json:"action,omitempty"`
}

func (m *WAFRuleResourceModel) ToAPIRequest() *APIWAFRule {
	return &APIWAFRule{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIWAFRuleSpec{
			Mode:   m.Mode.ValueString(),
			RuleID: m.RuleID.ValueString(),
			Action: m.Action.ValueString(),
		},
	}
}

func (m *WAFRuleResourceModel) FromAPIResponse(resp *APIWAFRule) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Mode = types.StringValue(resp.Spec.Mode)
	m.RuleID = types.StringValue(resp.Spec.RuleID)
	m.Action = types.StringValue(resp.Spec.Action)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
