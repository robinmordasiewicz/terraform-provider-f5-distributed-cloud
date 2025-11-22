// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package container_registry

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ContainerRegistryResourceModel describes the resource data model.
type ContainerRegistryResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// API structures

type APIContainerRegistry struct {
	Metadata   APIMetadata              `json:"metadata"`
	Spec       APIContainerRegistrySpec `json:"spec"`
	SystemMeta APISystemMetadata        `json:"system_metadata,omitempty"`
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

type APIContainerRegistrySpec struct {
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *ContainerRegistryResourceModel) ToAPIRequest() *APIContainerRegistry {
	spec := APIContainerRegistrySpec{}

	return &APIContainerRegistry{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *ContainerRegistryResourceModel) FromAPIResponse(resp *APIContainerRegistry) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
