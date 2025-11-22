// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cdn_cache_rule

import "github.com/hashicorp/terraform-plugin-framework/types"

type CDNCacheRuleResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	CacheTTL    types.Int64  `tfsdk:"cache_ttl"`
	PathMatch   types.String `tfsdk:"path_match"`
	ID          types.String `tfsdk:"id"`
}

type APICDNCacheRule struct {
	Metadata   APIMetadata         `json:"metadata"`
	Spec       APICDNCacheRuleSpec `json:"spec"`
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

type APICDNCacheRuleSpec struct {
	CacheTTL  int64  `json:"cache_ttl,omitempty"`
	PathMatch string `json:"path_match,omitempty"`
}

func (m *CDNCacheRuleResourceModel) ToAPIRequest() *APICDNCacheRule {
	spec := APICDNCacheRuleSpec{
		PathMatch: m.PathMatch.ValueString(),
	}
	if !m.CacheTTL.IsNull() {
		spec.CacheTTL = m.CacheTTL.ValueInt64()
	}
	return &APICDNCacheRule{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *CDNCacheRuleResourceModel) FromAPIResponse(resp *APICDNCacheRule) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.PathMatch = types.StringValue(resp.Spec.PathMatch)
	if resp.Spec.CacheTTL != 0 {
		m.CacheTTL = types.Int64Value(resp.Spec.CacheTTL)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
