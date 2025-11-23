// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package protocol_policer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &ProtocolPolicerDataSource{}

func NewProtocolPolicerDataSource() datasource.DataSource {
	return &ProtocolPolicerDataSource{}
}

type ProtocolPolicerDataSource struct {
	client *client.Client
}

type ProtocolPolicerDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Enabled     types.Bool   `tfsdk:"enabled"`
}

type APIProtocolPolicer struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Enabled bool `json:"enabled,omitempty"`
	} `json:"spec"`
}

func (d *ProtocolPolicerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocol_policer"
}

func (d *ProtocolPolicerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Protocol Policer.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_protocol_policer`" + ` data source retrieves information about an existing protocol policer.

## Example Usage

` + "```hcl" + `
data "f5_distributed_cloud_protocol_policer" "example" {
  name      = "my-protocol-policer"
  namespace = "my-namespace"
}

output "enabled" {
  value = data.f5_distributed_cloud_protocol_policer.example.enabled
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the protocol policer.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the protocol policer to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the protocol policer is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the protocol policer.",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the protocol policer is enabled.",
				Computed:    true,
			},
		},
	}
}

func (d *ProtocolPolicerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ProtocolPolicerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProtocolPolicerDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIProtocolPolicer
	path := fmt.Sprintf("/config/namespaces/%s/protocol_policers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Protocol Policer",
			fmt.Sprintf("Could not read protocol policer %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	data.Enabled = types.BoolValue(apiResp.Spec.Enabled)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
