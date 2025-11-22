// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package site_mesh_group

import "github.com/hashicorp/terraform-plugin-framework/types"

type SiteMeshGroupResourceModel struct {
	Name        types.String    `tfsdk:"name"`
	Namespace   types.String    `tfsdk:"namespace"`
	Description types.String    `tfsdk:"description"`
	Type        types.String    `tfsdk:"type"`
	VirtualSite []VirtualSiteRefModel `tfsdk:"virtual_site"`
	ID          types.String    `tfsdk:"id"`
}

type VirtualSiteRefModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type APISiteMeshGroup struct {
	Metadata   APIMetadata           `json:"metadata"`
	Spec       APISiteMeshGroupSpec  `json:"spec"`
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

type APISiteMeshGroupSpec struct {
	Type        string              `json:"type,omitempty"`
	VirtualSite []APIVirtualSiteRef `json:"virtual_site,omitempty"`
}

type APIVirtualSiteRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func (m *SiteMeshGroupResourceModel) ToAPIRequest() *APISiteMeshGroup {
	spec := APISiteMeshGroupSpec{
		Type: m.Type.ValueString(),
	}
	if len(m.VirtualSite) > 0 {
		spec.VirtualSite = make([]APIVirtualSiteRef, len(m.VirtualSite))
		for i, vs := range m.VirtualSite {
			spec.VirtualSite[i] = APIVirtualSiteRef{
				Name:      vs.Name.ValueString(),
				Namespace: vs.Namespace.ValueString(),
			}
		}
	}
	return &APISiteMeshGroup{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

func (m *SiteMeshGroupResourceModel) FromAPIResponse(resp *APISiteMeshGroup) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Type = types.StringValue(resp.Spec.Type)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
