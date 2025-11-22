// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_site

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &VirtualSiteDataSource{}

func NewVirtualSiteDataSource() datasource.DataSource {
	return &VirtualSiteDataSource{}
}

type VirtualSiteDataSource struct {
	client *client.Client
}

type VirtualSiteDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	SiteType    types.String `tfsdk:"site_type"`
}

type APIVirtualSite struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		SiteType string `json:"site_type,omitempty"`
	} `json:"spec"`
}

func (d *VirtualSiteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_site"
}

func (d *VirtualSiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Virtual Site.",
		MarkdownDescription: `
The ` + "`f5xc_virtual_site`" + ` data source retrieves information about an existing virtual site.

## Example Usage

` + "```hcl" + `
data "f5xc_virtual_site" "example" {
  name      = "my-virtual-site"
  namespace = "my-namespace"
}

output "site_type" {
  value = data.f5xc_virtual_site.example.site_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the virtual site.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the virtual site to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the virtual site is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the virtual site.",
				Computed:    true,
			},
			"site_type": schema.StringAttribute{
				Description: "Type of the virtual site (CUSTOMER_EDGE, REGIONAL_EDGE).",
				Computed:    true,
			},
		},
	}
}

func (d *VirtualSiteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *VirtualSiteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VirtualSiteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIVirtualSite
	path := fmt.Sprintf("/config/namespaces/%s/virtual_sites/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Virtual Site",
			fmt.Sprintf("Could not read virtual site %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.SiteType != "" {
		data.SiteType = types.StringValue(apiResp.Spec.SiteType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
