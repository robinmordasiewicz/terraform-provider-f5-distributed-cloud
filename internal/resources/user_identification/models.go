// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package user_identification

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UserIdentificationResourceModel describes the resource data model.
type UserIdentificationResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// API structures

// APIUserIdentification represents the API structure for user identification.
type APIUserIdentification struct {
	Metadata   APIMetadata               `json:"metadata"`
	Spec       APIUserIdentificationSpec `json:"spec"`
	SystemMeta APISystemMetadata         `json:"system_metadata,omitempty"`
}

// APIMetadata represents the metadata structure in API requests/responses.
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

// APIUserIdentificationSpec represents the spec structure for user identification.
type APIUserIdentificationSpec struct {
	// Spec fields can be extended based on actual API requirements
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *UserIdentificationResourceModel) ToAPIRequest() *APIUserIdentification {
	return &APIUserIdentification{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIUserIdentificationSpec{},
	}
}

// FromAPIResponse updates the model from an API response.
func (m *UserIdentificationResourceModel) FromAPIResponse(resp *APIUserIdentification) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
