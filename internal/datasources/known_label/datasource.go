// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package known_label

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &KnownLabelDataSource{}

func NewKnownLabelDataSource() datasource.DataSource {
	return &KnownLabelDataSource{}
}

type KnownLabelDataSource struct {
	client *client.Client
}

type KnownLabelDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	LabelKey    types.String `tfsdk:"label_key"`
}

type APIKnownLabel struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		LabelKey string `json:"label_key,omitempty"`
	} `json:"spec"`
}

func (d *KnownLabelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_known_label"
}

func (d *KnownLabelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Known Label.",
		MarkdownDescription: `
The ` + "`f5xc_known_label`" + ` data source retrieves information about an existing known label.

## Example Usage

` + "```hcl" + `
data "f5xc_known_label" "example" {
  name      = "my-known-label"
  namespace = "shared"
}

output "label_key" {
  value = data.f5xc_known_label.example.label_key
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the known label.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the known label to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the known label is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the known label.",
				Computed:    true,
			},
			"label_key": schema.StringAttribute{
				Description: "Key of the known label.",
				Computed:    true,
			},
		},
	}
}

func (d *KnownLabelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *KnownLabelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data KnownLabelDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIKnownLabel
	path := fmt.Sprintf("/config/namespaces/%s/known_labels/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Known Label",
			fmt.Sprintf("Could not read known label %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.LabelKey != "" {
		data.LabelKey = types.StringValue(apiResp.Spec.LabelKey)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
