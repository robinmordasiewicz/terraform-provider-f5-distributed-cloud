// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package gcp_vpc_site

import "github.com/hashicorp/terraform-plugin-framework/types"

type GCPVPCSiteResourceModel struct {
	Name                types.String `tfsdk:"name"`
	Namespace           types.String `tfsdk:"namespace"`
	Description         types.String `tfsdk:"description"`
	GCPRegion           types.String `tfsdk:"gcp_region"`
	ProjectID           types.String `tfsdk:"project_id"`
	NetworkName         types.String `tfsdk:"network_name"`
	MachineType         types.String `tfsdk:"machine_type"`
	CloudCredentialsRef types.String `tfsdk:"cloud_credentials_ref"`
	ID                  types.String `tfsdk:"id"`
}

type APIGCPVPCSite struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIGCPVPCSiteSpec `json:"spec"`
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

type APIGCPVPCSiteSpec struct {
	GCPRegion           string `json:"gcp_region,omitempty"`
	ProjectID           string `json:"project_id,omitempty"`
	NetworkName         string `json:"network_name,omitempty"`
	MachineType         string `json:"machine_type,omitempty"`
	CloudCredentialsRef string `json:"cloud_credentials_ref,omitempty"`
}

func (m *GCPVPCSiteResourceModel) ToAPIRequest() *APIGCPVPCSite {
	return &APIGCPVPCSite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIGCPVPCSiteSpec{
			GCPRegion:           m.GCPRegion.ValueString(),
			ProjectID:           m.ProjectID.ValueString(),
			NetworkName:         m.NetworkName.ValueString(),
			MachineType:         m.MachineType.ValueString(),
			CloudCredentialsRef: m.CloudCredentialsRef.ValueString(),
		},
	}
}

func (m *GCPVPCSiteResourceModel) FromAPIResponse(resp *APIGCPVPCSite) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.GCPRegion = types.StringValue(resp.Spec.GCPRegion)
	m.ProjectID = types.StringValue(resp.Spec.ProjectID)
	m.NetworkName = types.StringValue(resp.Spec.NetworkName)
	m.MachineType = types.StringValue(resp.Spec.MachineType)
	m.CloudCredentialsRef = types.StringValue(resp.Spec.CloudCredentialsRef)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
