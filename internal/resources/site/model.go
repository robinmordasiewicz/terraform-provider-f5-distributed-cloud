// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SiteResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

type APISite struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APISiteSpec        `json:"spec"`
	SystemMeta *APISystemMetadata `json:"system_metadata,omitempty"`
}

type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

type APISiteSpec struct {
}

func (m *SiteResourceModel) ToAPIRequest() *APISite {
	return &APISite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APISiteSpec{},
	}
}

func (m *SiteResourceModel) FromAPIResponse(resp *APISite) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	if resp.Metadata.Description != "" {
		m.Description = types.StringValue(resp.Metadata.Description)
	}
	if resp.SystemMeta != nil && resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
