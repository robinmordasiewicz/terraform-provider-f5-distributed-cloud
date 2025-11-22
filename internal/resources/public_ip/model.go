// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package public_ip

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PublicIPResourceModel describes the resource data model.
type PublicIPResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// API structures

type APIPublicIP struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIPublicIPSpec   `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
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

type APIPublicIPSpec struct {
	// Spec fields can be extended as needed
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *PublicIPResourceModel) ToAPIRequest() *APIPublicIP {
	return &APIPublicIP{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIPublicIPSpec{},
	}
}

// FromAPIResponse updates the model from an API response.
func (m *PublicIPResourceModel) FromAPIResponse(resp *APIPublicIP) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
