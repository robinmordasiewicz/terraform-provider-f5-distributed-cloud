// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_mesh

import "github.com/hashicorp/terraform-plugin-framework/types"

type ServiceMeshResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	MeshType    types.String `tfsdk:"mesh_type"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	ID          types.String `tfsdk:"id"`
}

type APIServiceMesh struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIServiceMeshSpec `json:"spec"`
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

type APIServiceMeshSpec struct {
	MeshType string `json:"mesh_type,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
}

func (m *ServiceMeshResourceModel) ToAPIRequest() *APIServiceMesh {
	return &APIServiceMesh{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIServiceMeshSpec{
			MeshType: m.MeshType.ValueString(),
			Enabled:  m.Enabled.ValueBool(),
		},
	}
}

func (m *ServiceMeshResourceModel) FromAPIResponse(resp *APIServiceMesh) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.MeshType = types.StringValue(resp.Spec.MeshType)
	m.Enabled = types.BoolValue(resp.Spec.Enabled)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
