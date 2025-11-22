// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package endpoint

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// EndpointResourceModel describes the resource data model.
type EndpointResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

// API structures

type APIEndpoint struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIEndpointSpec   `json:"spec"`
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

type APIEndpointSpec struct {
	// Empty spec for basic endpoint
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *EndpointResourceModel) ToAPIRequest() *APIEndpoint {
	return &APIEndpoint{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIEndpointSpec{},
	}
}

// FromAPIResponse updates the model from an API response.
func (m *EndpointResourceModel) FromAPIResponse(resp *APIEndpoint) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
