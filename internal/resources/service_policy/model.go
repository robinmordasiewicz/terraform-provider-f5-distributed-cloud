// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServicePolicyResourceModel describes the resource data model.
type ServicePolicyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Policy rules
	Rules []RuleModel `tfsdk:"rules"`

	// Policy action when no rules match
	DefaultAction types.String `tfsdk:"default_action"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// RuleModel represents a service policy rule.
type RuleModel struct {
	Name   types.String `tfsdk:"name"`
	Action types.String `tfsdk:"action"`

	// Match conditions
	MatchSources      []MatchSourceModel `tfsdk:"match_sources"`
	MatchDestinations []MatchDestModel   `tfsdk:"match_destinations"`
}

// MatchSourceModel represents source matching criteria.
type MatchSourceModel struct {
	Prefix     types.String `tfsdk:"prefix"`
	PrefixList types.String `tfsdk:"prefix_list"`
	ASN        types.Int64  `tfsdk:"asn"`
	Namespace  types.String `tfsdk:"namespace"`
}

// MatchDestModel represents destination matching criteria.
type MatchDestModel struct {
	Prefix   types.String `tfsdk:"prefix"`
	Port     types.Int64  `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}

// API structures

type APIServicePolicy struct {
	Metadata   APIMetadata          `json:"metadata"`
	Spec       APIServicePolicySpec `json:"spec"`
	SystemMeta APISystemMetadata    `json:"system_metadata,omitempty"`
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

type APIServicePolicySpec struct {
	Rules         []APIRule         `json:"rules,omitempty"`
	DefaultAction *APIDefaultAction `json:"default_action,omitempty"`
}

type APIRule struct {
	Metadata APIRuleMetadata `json:"metadata,omitempty"`
	Spec     APIRuleSpec     `json:"spec,omitempty"`
}

type APIRuleMetadata struct {
	Name string `json:"name,omitempty"`
}

type APIRuleSpec struct {
	Action string `json:"action,omitempty"`
}

type APIDefaultAction struct {
	AllowAll *struct{} `json:"allow,omitempty"`
	DenyAll  *struct{} `json:"deny,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *ServicePolicyResourceModel) ToAPIRequest() *APIServicePolicy {
	spec := APIServicePolicySpec{}

	if len(m.Rules) > 0 {
		rules := make([]APIRule, len(m.Rules))
		for i, r := range m.Rules {
			rules[i] = APIRule{
				Metadata: APIRuleMetadata{
					Name: r.Name.ValueString(),
				},
				Spec: APIRuleSpec{
					Action: r.Action.ValueString(),
				},
			}
		}
		spec.Rules = rules
	}

	if !m.DefaultAction.IsNull() && !m.DefaultAction.IsUnknown() {
		action := m.DefaultAction.ValueString()
		spec.DefaultAction = &APIDefaultAction{}
		if action == "allow" {
			spec.DefaultAction.AllowAll = &struct{}{}
		} else if action == "deny" {
			spec.DefaultAction.DenyAll = &struct{}{}
		}
	}

	return &APIServicePolicy{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *ServicePolicyResourceModel) FromAPIResponse(resp *APIServicePolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if len(resp.Spec.Rules) > 0 {
		rules := make([]RuleModel, len(resp.Spec.Rules))
		for i, r := range resp.Spec.Rules {
			rules[i] = RuleModel{
				Name:   types.StringValue(r.Metadata.Name),
				Action: types.StringValue(r.Spec.Action),
			}
		}
		m.Rules = rules
	}

	if resp.Spec.DefaultAction != nil {
		if resp.Spec.DefaultAction.AllowAll != nil {
			m.DefaultAction = types.StringValue("allow")
		} else if resp.Spec.DefaultAction.DenyAll != nil {
			m.DefaultAction = types.StringValue("deny")
		}
	}

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
