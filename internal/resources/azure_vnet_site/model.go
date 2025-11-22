// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package azure_vnet_site

import "github.com/hashicorp/terraform-plugin-framework/types"

type AzureVNETSiteResourceModel struct {
	Name                types.String `tfsdk:"name"`
	Namespace           types.String `tfsdk:"namespace"`
	Description         types.String `tfsdk:"description"`
	AzureRegion         types.String `tfsdk:"azure_region"`
	ResourceGroup       types.String `tfsdk:"resource_group"`
	VNETName            types.String `tfsdk:"vnet_name"`
	MachineType         types.String `tfsdk:"machine_type"`
	CloudCredentialsRef types.String `tfsdk:"cloud_credentials_ref"`
	ID                  types.String `tfsdk:"id"`
}

type APIAzureVNETSite struct {
	Metadata   APIMetadata          `json:"metadata"`
	Spec       APIAzureVNETSiteSpec `json:"spec"`
	SystemMeta APISystemMetadata    `json:"system_metadata,omitempty"`
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

type APIAzureVNETSiteSpec struct {
	AzureRegion         string `json:"azure_region,omitempty"`
	ResourceGroup       string `json:"resource_group,omitempty"`
	VNETName            string `json:"vnet_name,omitempty"`
	MachineType         string `json:"machine_type,omitempty"`
	CloudCredentialsRef string `json:"cloud_credentials_ref,omitempty"`
}

func (m *AzureVNETSiteResourceModel) ToAPIRequest() *APIAzureVNETSite {
	return &APIAzureVNETSite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIAzureVNETSiteSpec{
			AzureRegion:         m.AzureRegion.ValueString(),
			ResourceGroup:       m.ResourceGroup.ValueString(),
			VNETName:            m.VNETName.ValueString(),
			MachineType:         m.MachineType.ValueString(),
			CloudCredentialsRef: m.CloudCredentialsRef.ValueString(),
		},
	}
}

func (m *AzureVNETSiteResourceModel) FromAPIResponse(resp *APIAzureVNETSite) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.AzureRegion = types.StringValue(resp.Spec.AzureRegion)
	m.ResourceGroup = types.StringValue(resp.Spec.ResourceGroup)
	m.VNETName = types.StringValue(resp.Spec.VNETName)
	m.MachineType = types.StringValue(resp.Spec.MachineType)
	m.CloudCredentialsRef = types.StringValue(resp.Spec.CloudCredentialsRef)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
