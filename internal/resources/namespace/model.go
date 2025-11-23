// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NamespaceResourceModel describes the resource data model.
type NamespaceResourceModel struct {
	// User-configurable fields
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`

	// Computed fields (read-only from API)
	ID types.String `tfsdk:"id"`

	// Timeouts configuration
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

// APINamespace represents the namespace structure for API requests/responses.
type APINamespace struct {
	Metadata   APINamespaceMetadata `json:"metadata"`
	Spec       APINamespaceSpec     `json:"spec,omitempty"`
	SystemMeta APISystemMetadata    `json:"system_metadata,omitempty"`
}

// APINamespaceMetadata represents the metadata for a namespace.
type APINamespaceMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// APINamespaceSpec represents the spec for a namespace (currently empty).
type APINamespaceSpec struct {
	// Namespace spec is typically empty for basic namespaces
}

// APISystemMetadata represents system-managed metadata from the API.
type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
	ModificationTime  string `json:"modification_timestamp,omitempty"`
	CreatorClass      string `json:"creator_class,omitempty"`
	CreatorID         string `json:"creator_id,omitempty"`
}

// APINamespaceList represents a list of namespaces from the API.
type APINamespaceList struct {
	Items []APINamespace `json:"items"`
}

// ToAPIRequest converts the Terraform model to an API request structure.
func (m *NamespaceResourceModel) ToAPIRequest() *APINamespace {
	return &APINamespace{
		Metadata: APINamespaceMetadata{
			Name:        m.Name.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APINamespaceSpec{},
	}
}

// FromAPIResponse updates the Terraform model from an API response.
func (m *NamespaceResourceModel) FromAPIResponse(resp *APINamespace) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Description = types.StringValue(resp.Metadata.Description)

	// Set ID from the UID if available, otherwise use name
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Name)
	}
}
