// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package voltstack_site

import "github.com/hashicorp/terraform-plugin-framework/types"

type VoltstackSiteResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Namespace       types.String `tfsdk:"namespace"`
	Description     types.String `tfsdk:"description"`
	MasterNodes     types.Int64  `tfsdk:"master_nodes"`
	WorkerNodes     types.Int64  `tfsdk:"worker_nodes"`
	SiteType        types.String `tfsdk:"site_type"`
	VolterraRegion  types.String `tfsdk:"volterra_region"`
	ID              types.String `tfsdk:"id"`
}

type APIVoltstackSite struct {
	Metadata   APIMetadata           `json:"metadata"`
	Spec       APIVoltstackSiteSpec  `json:"spec"`
	SystemMeta APISystemMetadata     `json:"system_metadata,omitempty"`
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

type APIVoltstackSiteSpec struct {
	MasterNodes    int64  `json:"master_nodes,omitempty"`
	WorkerNodes    int64  `json:"worker_nodes,omitempty"`
	SiteType       string `json:"site_type,omitempty"`
	VolterraRegion string `json:"volterra_region,omitempty"`
}

func (m *VoltstackSiteResourceModel) ToAPIRequest() *APIVoltstackSite {
	return &APIVoltstackSite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIVoltstackSiteSpec{
			MasterNodes:    m.MasterNodes.ValueInt64(),
			WorkerNodes:    m.WorkerNodes.ValueInt64(),
			SiteType:       m.SiteType.ValueString(),
			VolterraRegion: m.VolterraRegion.ValueString(),
		},
	}
}

func (m *VoltstackSiteResourceModel) FromAPIResponse(resp *APIVoltstackSite) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.MasterNodes = types.Int64Value(resp.Spec.MasterNodes)
	m.WorkerNodes = types.Int64Value(resp.Spec.WorkerNodes)
	m.SiteType = types.StringValue(resp.Spec.SiteType)
	m.VolterraRegion = types.StringValue(resp.Spec.VolterraRegion)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
