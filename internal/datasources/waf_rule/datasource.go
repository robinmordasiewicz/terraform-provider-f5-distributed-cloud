// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package waf_rule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &WAFRuleDataSource{}

func NewWAFRuleDataSource() datasource.DataSource {
	return &WAFRuleDataSource{}
}

type WAFRuleDataSource struct {
	client *client.Client
}

type WAFRuleDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Mode        types.String `tfsdk:"mode"`
}

type APIWAFRule struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Mode string `json:"mode,omitempty"`
	} `json:"spec"`
}

func (d *WAFRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_waf_rule"
}

func (d *WAFRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud WAF Rule.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_waf_rule`" + ` data source retrieves information about an existing WAF rule.

## Example Usage

` + "```hcl" + `
data "f5_distributed_cloud_waf_rule" "example" {
  name      = "my-waf-rule"
  namespace = "my-namespace"
}

output "waf_mode" {
  value = data.f5_distributed_cloud_waf_rule.example.mode
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the WAF rule.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the WAF rule to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the WAF rule is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the WAF rule.",
				Computed:    true,
			},
			"mode": schema.StringAttribute{
				Description: "Mode of the WAF rule (BLOCK, DETECT).",
				Computed:    true,
			},
		},
	}
}

func (d *WAFRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *WAFRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data WAFRuleDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIWAFRule
	path := fmt.Sprintf("/config/namespaces/%s/waf_rules/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading WAF Rule",
			fmt.Sprintf("Could not read WAF rule %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.Mode != "" {
		data.Mode = types.StringValue(apiResp.Spec.Mode)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
