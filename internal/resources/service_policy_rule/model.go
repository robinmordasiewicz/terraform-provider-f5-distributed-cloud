// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_policy_rule

import "github.com/hashicorp/terraform-plugin-framework/types"

type ServicePolicyRuleResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

type APIServicePolicyRule struct {
	Metadata   APIMetadata                `json:"metadata"`
	Spec       APIServicePolicyRuleSpec   `json:"spec"`
	SystemMeta APISystemMetadata          `json:"system_metadata,omitempty"`
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

type APIServicePolicyRuleSpec struct {
}

func (m *ServicePolicyRuleResourceModel) ToAPIRequest() *APIServicePolicyRule {
	spec := APIServicePolicyRuleSpec{}
	return &APIServicePolicyRule{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *ServicePolicyRuleResourceModel) FromAPIResponse(resp *APIServicePolicyRule) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
