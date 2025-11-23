// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tunnel

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &TunnelDataSource{}

func NewTunnelDataSource() datasource.DataSource {
	return &TunnelDataSource{}
}

type TunnelDataSource struct {
	client *client.Client
}

type TunnelDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	TunnelType  types.String `tfsdk:"tunnel_type"`
}

type APITunnel struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		TunnelType string `json:"tunnel_type,omitempty"`
	} `json:"spec"`
}

func (d *TunnelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnel"
}

func (d *TunnelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Tunnel.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_tunnel`" + ` data source retrieves information about an existing tunnel.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_tunnel" "example" {
  name      = "my-tunnel"
  namespace = "my-namespace"
}

output "tunnel_type" {
  value = data.f5distributedcloud_tunnel.example.tunnel_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the tunnel.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the tunnel to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the tunnel is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the tunnel.",
				Computed:    true,
			},
			"tunnel_type": schema.StringAttribute{
				Description: "Type of tunnel (IPSEC, SSL).",
				Computed:    true,
			},
		},
	}
}

func (d *TunnelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TunnelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TunnelDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APITunnel
	path := fmt.Sprintf("/config/namespaces/%s/tunnels/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Tunnel",
			fmt.Sprintf("Could not read tunnel %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.TunnelType != "" {
		data.TunnelType = types.StringValue(apiResp.Spec.TunnelType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
