// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TokenResourceModel describes the resource data model.
type TokenResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// API structures

// APIToken represents the API request/response structure for a Token.
type APIToken struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APITokenSpec      `json:"spec,omitempty"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
}

// APIMetadata represents common metadata fields.
type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

// APISystemMetadata represents system-generated metadata.
type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
}

// APITokenSpec represents the specification for a Token.
type APITokenSpec struct {
	// Token spec fields can be extended as needed
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *TokenResourceModel) ToAPIRequest() *APIToken {
	return &APIToken{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APITokenSpec{},
	}
}

// FromAPIResponse updates the model from an API response.
func (m *TokenResourceModel) FromAPIResponse(resp *APIToken) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)

	if resp.Metadata.Description != "" {
		m.Description = types.StringValue(resp.Metadata.Description)
	} else {
		m.Description = types.StringNull()
	}

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
