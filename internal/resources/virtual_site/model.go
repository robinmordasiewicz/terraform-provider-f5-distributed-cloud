// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VirtualSiteResourceModel describes the resource data model.
type VirtualSiteResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Site type
	SiteType types.String `tfsdk:"site_type"`

	// Site selector
	SiteSelector *SiteSelectorModel `tfsdk:"site_selector"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// SiteSelectorModel represents the site selector configuration.
type SiteSelectorModel struct {
	Expressions []types.String `tfsdk:"expressions"`
}

// API structures

type APIVirtualSite struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIVirtualSiteSpec `json:"spec"`
	SystemMeta APISystemMetadata  `json:"system_metadata,omitempty"`
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

type APIVirtualSiteSpec struct {
	SiteType     string           `json:"site_type,omitempty"`
	SiteSelector *APISiteSelector `json:"site_selector,omitempty"`
}

type APISiteSelector struct {
	Expressions []string `json:"expressions,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *VirtualSiteResourceModel) ToAPIRequest() *APIVirtualSite {
	spec := APIVirtualSiteSpec{}

	if !m.SiteType.IsNull() && !m.SiteType.IsUnknown() {
		spec.SiteType = m.SiteType.ValueString()
	}

	if m.SiteSelector != nil && len(m.SiteSelector.Expressions) > 0 {
		exprs := make([]string, len(m.SiteSelector.Expressions))
		for i, e := range m.SiteSelector.Expressions {
			exprs[i] = e.ValueString()
		}
		spec.SiteSelector = &APISiteSelector{
			Expressions: exprs,
		}
	}

	return &APIVirtualSite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *VirtualSiteResourceModel) FromAPIResponse(resp *APIVirtualSite) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.Spec.SiteType != "" {
		m.SiteType = types.StringValue(resp.Spec.SiteType)
	}

	if resp.Spec.SiteSelector != nil && len(resp.Spec.SiteSelector.Expressions) > 0 {
		exprs := make([]types.String, len(resp.Spec.SiteSelector.Expressions))
		for i, e := range resp.Spec.SiteSelector.Expressions {
			exprs[i] = types.StringValue(e)
		}
		m.SiteSelector = &SiteSelectorModel{
			Expressions: exprs,
		}
	}

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
