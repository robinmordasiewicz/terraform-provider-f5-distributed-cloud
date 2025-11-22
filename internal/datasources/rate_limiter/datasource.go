// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package rate_limiter

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &RateLimiterDataSource{}

func NewRateLimiterDataSource() datasource.DataSource {
	return &RateLimiterDataSource{}
}

type RateLimiterDataSource struct {
	client *client.Client
}

type RateLimiterDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	TotalNumber types.Int64  `tfsdk:"total_number"`
	Unit        types.String `tfsdk:"unit"`
	BurstMultiplier types.Int64 `tfsdk:"burst_multiplier"`
}

type APIRateLimiter struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		TotalNumber     int    `json:"total_number,omitempty"`
		Unit            string `json:"unit,omitempty"`
		BurstMultiplier int    `json:"burst_multiplier,omitempty"`
	} `json:"spec"`
}

func (d *RateLimiterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rate_limiter"
}

func (d *RateLimiterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Rate Limiter.",
		MarkdownDescription: `
The ` + "`f5xc_rate_limiter`" + ` data source retrieves information about an existing rate limiter configuration.

## Example Usage

` + "```hcl" + `
data "f5xc_rate_limiter" "example" {
  name      = "my-rate-limiter"
  namespace = "my-namespace"
}

output "rate_limit" {
  value = "${data.f5xc_rate_limiter.example.total_number} per ${data.f5xc_rate_limiter.example.unit}"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the rate limiter.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the rate limiter to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the rate limiter is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the rate limiter.",
				Computed:    true,
			},
			"total_number": schema.Int64Attribute{
				Description: "Maximum number of requests allowed in the time unit.",
				Computed:    true,
			},
			"unit": schema.StringAttribute{
				Description: "Time unit for rate limiting (SECOND, MINUTE).",
				Computed:    true,
			},
			"burst_multiplier": schema.Int64Attribute{
				Description: "Multiplier for burst capacity.",
				Computed:    true,
			},
		},
	}
}

func (d *RateLimiterDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *RateLimiterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RateLimiterDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIRateLimiter
	path := fmt.Sprintf("/config/namespaces/%s/rate_limiters/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Rate Limiter",
			fmt.Sprintf("Could not read rate limiter %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	data.TotalNumber = types.Int64Value(int64(apiResp.Spec.TotalNumber))
	if apiResp.Spec.Unit != "" {
		data.Unit = types.StringValue(apiResp.Spec.Unit)
	}
	data.BurstMultiplier = types.Int64Value(int64(apiResp.Spec.BurstMultiplier))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
