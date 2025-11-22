// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package ip_prefix_set

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &IPPrefixSetDataSource{}

func NewIPPrefixSetDataSource() datasource.DataSource {
	return &IPPrefixSetDataSource{}
}

type IPPrefixSetDataSource struct {
	client *client.Client
}

type IPPrefixSetDataSourceModel struct {
	ID          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Namespace   types.String   `tfsdk:"namespace"`
	Description types.String   `tfsdk:"description"`
	Prefixes    []types.String `tfsdk:"prefixes"`
}

type APIIPPrefixSet struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Prefixes []string `json:"prefixes,omitempty"`
	} `json:"spec"`
}

func (d *IPPrefixSetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_prefix_set"
}

func (d *IPPrefixSetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud IP Prefix Set.",
		MarkdownDescription: `
The ` + "`f5xc_ip_prefix_set`" + ` data source retrieves information about an existing IP prefix set.

## Example Usage

` + "```hcl" + `
data "f5xc_ip_prefix_set" "example" {
  name      = "my-prefix-set"
  namespace = "my-namespace"
}

output "prefixes" {
  value = data.f5xc_ip_prefix_set.example.prefixes
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the IP prefix set.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the IP prefix set to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the IP prefix set is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the IP prefix set.",
				Computed:    true,
			},
			"prefixes": schema.ListAttribute{
				Description: "List of IP prefixes in CIDR notation.",
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *IPPrefixSetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *IPPrefixSetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IPPrefixSetDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIIPPrefixSet
	path := fmt.Sprintf("/config/namespaces/%s/ip_prefix_sets/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading IP Prefix Set",
			fmt.Sprintf("Could not read IP prefix set %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if len(apiResp.Spec.Prefixes) > 0 {
		data.Prefixes = make([]types.String, len(apiResp.Spec.Prefixes))
		for i, prefix := range apiResp.Spec.Prefixes {
			data.Prefixes[i] = types.StringValue(prefix)
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
