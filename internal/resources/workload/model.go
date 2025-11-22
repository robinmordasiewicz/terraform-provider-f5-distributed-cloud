// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package workload

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// WorkloadResourceModel describes the resource data model.
type WorkloadResourceModel struct {
	// User-configurable fields
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// APIWorkload represents the workload structure for API requests/responses.
type APIWorkload struct {
	Metadata   APIWorkloadMetadata `json:"metadata"`
	Spec       APIWorkloadSpec     `json:"spec,omitempty"`
	SystemMeta APISystemMetadata   `json:"system_metadata,omitempty"`
}

// APIWorkloadMetadata represents the metadata for a workload.
type APIWorkloadMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

// APIWorkloadSpec represents the spec for a workload.
type APIWorkloadSpec struct {
	// Workload spec - can be extended as needed
}

// APISystemMetadata represents system-managed metadata from the API.
type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
	ModificationTime  string `json:"modification_timestamp,omitempty"`
	CreatorClass      string `json:"creator_class,omitempty"`
	CreatorID         string `json:"creator_id,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request structure.
func (m *WorkloadResourceModel) ToAPIRequest() *APIWorkload {
	return &APIWorkload{
		Metadata: APIWorkloadMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIWorkloadSpec{},
	}
}

// FromAPIResponse updates the Terraform model from an API response.
func (m *WorkloadResourceModel) FromAPIResponse(resp *APIWorkload) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	// Set ID from the UID if available, otherwise use namespace/name
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
