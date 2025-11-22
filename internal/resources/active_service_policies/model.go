// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package active_service_policies

import "github.com/hashicorp/terraform-plugin-framework/types"

type ActiveServicePoliciesResourceModel struct {
	Name        types.String     `tfsdk:"name"`
	Namespace   types.String     `tfsdk:"namespace"`
	Description types.String     `tfsdk:"description"`
	Policies    []PolicyRefModel `tfsdk:"policies"`
	ID          types.String     `tfsdk:"id"`
}

type PolicyRefModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type APIActiveServicePolicies struct {
	Metadata   APIMetadata                  `json:"metadata"`
	Spec       APIActiveServicePoliciesSpec `json:"spec"`
	SystemMeta APISystemMetadata            `json:"system_metadata,omitempty"`
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

type APIActiveServicePoliciesSpec struct {
	Policies []APIPolicyRef `json:"policies,omitempty"`
}

type APIPolicyRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func (m *ActiveServicePoliciesResourceModel) ToAPIRequest() *APIActiveServicePolicies {
	spec := APIActiveServicePoliciesSpec{}
	if len(m.Policies) > 0 {
		spec.Policies = make([]APIPolicyRef, len(m.Policies))
		for i, policy := range m.Policies {
			spec.Policies[i] = APIPolicyRef{
				Name:      policy.Name.ValueString(),
				Namespace: policy.Namespace.ValueString(),
			}
		}
	}
	return &APIActiveServicePolicies{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *ActiveServicePoliciesResourceModel) FromAPIResponse(resp *APIActiveServicePolicies) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
