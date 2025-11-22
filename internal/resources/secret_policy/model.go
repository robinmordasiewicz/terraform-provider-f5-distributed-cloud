// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package secret_policy

import "github.com/hashicorp/terraform-plugin-framework/types"

type SecretPolicyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	PolicyType  types.String `tfsdk:"policy_type"`
	Rules       types.String `tfsdk:"rules"`
	ID          types.String `tfsdk:"id"`
}

type APISecretPolicy struct {
	Metadata   APIMetadata         `json:"metadata"`
	Spec       APISecretPolicySpec `json:"spec"`
	SystemMeta APISystemMetadata   `json:"system_metadata,omitempty"`
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

type APISecretPolicySpec struct {
	PolicyType string `json:"policy_type,omitempty"`
	Rules      string `json:"rules,omitempty"`
}

func (m *SecretPolicyResourceModel) ToAPIRequest() *APISecretPolicy {
	return &APISecretPolicy{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APISecretPolicySpec{
			PolicyType: m.PolicyType.ValueString(),
			Rules:      m.Rules.ValueString(),
		},
	}
}

func (m *SecretPolicyResourceModel) FromAPIResponse(resp *APISecretPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.PolicyType = types.StringValue(resp.Spec.PolicyType)
	m.Rules = types.StringValue(resp.Spec.Rules)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
