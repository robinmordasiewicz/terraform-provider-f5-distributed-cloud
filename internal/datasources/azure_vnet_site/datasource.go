// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package azure_vnet_site

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &AzureVNETSiteDataSource{}

func NewAzureVNETSiteDataSource() datasource.DataSource {
	return &AzureVNETSiteDataSource{}
}

type AzureVNETSiteDataSource struct {
	client *client.Client
}

type AzureVNETSiteDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	SiteState   types.String `tfsdk:"site_state"`
	AzureRegion types.String `tfsdk:"azure_region"`
}

type APIAzureVNETSite struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		AzureRegion string `json:"azure_region,omitempty"`
	} `json:"spec"`
	Status struct {
		SiteState string `json:"site_state,omitempty"`
	} `json:"status"`
}

func (d *AzureVNETSiteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_azure_vnet_site"
}

func (d *AzureVNETSiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Azure VNET Site.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_azure_vnet_site`" + ` data source retrieves information about an existing Azure VNET site.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_azure_vnet_site" "example" {
  name      = "my-azure-site"
  namespace = "system"
}

output "site_state" {
  value = data.f5distributedcloud_azure_vnet_site.example.site_state
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the Azure VNET site.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Azure VNET site to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Azure VNET site is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the Azure VNET site.",
				Computed:    true,
			},
			"site_state": schema.StringAttribute{
				Description: "State of the site (ONLINE, OFFLINE, etc.).",
				Computed:    true,
			},
			"azure_region": schema.StringAttribute{
				Description: "Azure region where the site is deployed.",
				Computed:    true,
			},
		},
	}
}

func (d *AzureVNETSiteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData))
		return
	}
	d.client = c
}

func (d *AzureVNETSiteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AzureVNETSiteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIAzureVNETSite
	path := fmt.Sprintf("/config/namespaces/%s/azure_vnet_sites/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Azure VNET Site",
			fmt.Sprintf("Could not read Azure VNET site %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	data.Name = types.StringValue(apiResp.Metadata.Name)
	data.Namespace = types.StringValue(apiResp.Metadata.Namespace)
	if apiResp.Metadata.UID != "" {
		data.ID = types.StringValue(apiResp.Metadata.UID)
	} else {
		data.ID = types.StringValue(fmt.Sprintf("%s/%s", apiResp.Metadata.Namespace, apiResp.Metadata.Name))
	}
	if apiResp.Metadata.Description != "" {
		data.Description = types.StringValue(apiResp.Metadata.Description)
	}
	if apiResp.Status.SiteState != "" {
		data.SiteState = types.StringValue(apiResp.Status.SiteState)
	}
	if apiResp.Spec.AzureRegion != "" {
		data.AzureRegion = types.StringValue(apiResp.Spec.AzureRegion)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
