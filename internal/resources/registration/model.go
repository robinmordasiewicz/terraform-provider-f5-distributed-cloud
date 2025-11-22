// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package registration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RegistrationResourceModel describes the resource data model.
type RegistrationResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// API structures

// APIRegistration represents the API request/response structure.
type APIRegistration struct {
	Metadata   APIMetadata         `json:"metadata"`
	Spec       APIRegistrationSpec `json:"spec"`
	SystemMeta APISystemMetadata   `json:"system_metadata,omitempty"`
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

// APIRegistrationSpec represents the registration specification.
type APIRegistrationSpec struct {
	// Add registration-specific spec fields here as needed
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *RegistrationResourceModel) ToAPIRequest() *APIRegistration {
	return &APIRegistration{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIRegistrationSpec{},
	}
}

// FromAPIResponse updates the model from an API response.
func (m *RegistrationResourceModel) FromAPIResponse(resp *APIRegistration) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
