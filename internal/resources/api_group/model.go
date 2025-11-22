// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_group

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// APIGroupResourceModel describes the resource data model.
type APIGroupResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// APIAPIGroup represents the API request/response structure.
type APIAPIGroup struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIAPIGroupSpec    `json:"spec"`
	SystemMeta *APISystemMetadata `json:"system_metadata,omitempty"`
}

// APIMetadata represents the metadata section.
type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

// APISystemMetadata represents system-generated metadata.
type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

// APIAPIGroupSpec represents the API Group specification.
type APIAPIGroupSpec struct {
}

// ToAPIRequest converts the Terraform model to API request format.
func (m *APIGroupResourceModel) ToAPIRequest() *APIAPIGroup {
	apiReq := &APIAPIGroup{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIAPIGroupSpec{},
	}

	return apiReq
}

// FromAPIResponse updates the Terraform model from API response.
func (m *APIGroupResourceModel) FromAPIResponse(resp *APIAPIGroup) {
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
