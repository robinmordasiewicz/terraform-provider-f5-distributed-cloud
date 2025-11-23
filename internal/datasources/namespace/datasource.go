// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &NamespaceDataSource{}

func NewNamespaceDataSource() datasource.DataSource {
	return &NamespaceDataSource{}
}

type NamespaceDataSource struct {
	client *client.Client
}

type NamespaceDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Labels      types.Map    `tfsdk:"labels"`
}

type APINamespace struct {
	Metadata struct {
		Name        string            `json:"name"`
		Description string            `json:"description,omitempty"`
		Labels      map[string]string `json:"labels,omitempty"`
		UID         string            `json:"uid,omitempty"`
	} `json:"metadata"`
}

func (d *NamespaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_namespace"
}

func (d *NamespaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Namespace.",
		MarkdownDescription: `
The ` + "`f5xc_namespace`" + ` data source retrieves information about an existing namespace.

## Example Usage

` + "```hcl" + `
data "f5xc_namespace" "example" {
  name = "my-namespace"
}

output "namespace_id" {
  value = data.f5xc_namespace.example.id
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the namespace.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the namespace to look up.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the namespace.",
				Computed:    true,
			},
			"labels": schema.MapAttribute{
				Description: "Labels applied to the namespace.",
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *NamespaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NamespaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NamespaceDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APINamespace
	path := fmt.Sprintf("/web/namespaces/%s", data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Namespace",
			fmt.Sprintf("Could not read namespace %s: %s", data.Name.ValueString(), err))
		return
	}

	// Update model from API response
	data.Name = types.StringValue(apiResp.Metadata.Name)
	if apiResp.Metadata.UID != "" {
		data.ID = types.StringValue(apiResp.Metadata.UID)
	} else {
		data.ID = types.StringValue(apiResp.Metadata.Name)
	}
	if apiResp.Metadata.Description != "" {
		data.Description = types.StringValue(apiResp.Metadata.Description)
	}

	// Convert labels
	if len(apiResp.Metadata.Labels) > 0 {
		labels := make(map[string]string)
		for k, v := range apiResp.Metadata.Labels {
			labels[k] = v
		}
		labelsMap, diags := types.MapValueFrom(ctx, types.StringType, labels)
		resp.Diagnostics.Append(diags...)
		data.Labels = labelsMap
	} else {
		data.Labels = types.MapNull(types.StringType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
