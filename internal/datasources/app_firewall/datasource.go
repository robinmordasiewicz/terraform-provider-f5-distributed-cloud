// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_firewall

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &AppFirewallDataSource{}

func NewAppFirewallDataSource() datasource.DataSource {
	return &AppFirewallDataSource{}
}

type AppFirewallDataSource struct {
	client *client.Client
}

type AppFirewallDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Mode        types.String `tfsdk:"mode"`
}

type APIAppFirewall struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace,omitempty"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Mode string `json:"mode,omitempty"`
	} `json:"spec"`
}

func (d *AppFirewallDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_firewall"
}

func (d *AppFirewallDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Application Firewall.",
		MarkdownDescription: `
The ` + "`f5xc_app_firewall`" + ` data source retrieves information about an existing Application Firewall.

## Example Usage

` + "```hcl" + `
data "f5xc_app_firewall" "example" {
  name      = "my-waf"
  namespace = "my-namespace"
}

output "waf_mode" {
  value = data.f5xc_app_firewall.example.mode
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the Application Firewall.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Application Firewall to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace of the Application Firewall.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the Application Firewall.",
				Computed:    true,
			},
			"mode": schema.StringAttribute{
				Description: "Mode of the Application Firewall (e.g., BLOCK, DETECT).",
				Computed:    true,
			},
		},
	}
}

func (d *AppFirewallDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *AppFirewallDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppFirewallDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIAppFirewall
	path := fmt.Sprintf("/config/namespaces/%s/app_firewalls/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Application Firewall",
			fmt.Sprintf("Could not read Application Firewall %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	// Update model from API response
	data.Name = types.StringValue(apiResp.Metadata.Name)
	data.Namespace = types.StringValue(apiResp.Metadata.Namespace)
	if apiResp.Metadata.UID != "" {
		data.ID = types.StringValue(apiResp.Metadata.UID)
	} else {
		data.ID = types.StringValue(apiResp.Metadata.Namespace + "/" + apiResp.Metadata.Name)
	}
	if apiResp.Metadata.Description != "" {
		data.Description = types.StringValue(apiResp.Metadata.Description)
	}
	if apiResp.Spec.Mode != "" {
		data.Mode = types.StringValue(apiResp.Spec.Mode)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
