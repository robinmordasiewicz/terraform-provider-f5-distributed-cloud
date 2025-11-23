// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package healthcheck

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &HealthcheckDataSource{}

func NewHealthcheckDataSource() datasource.DataSource {
	return &HealthcheckDataSource{}
}

type HealthcheckDataSource struct {
	client *client.Client
}

type HealthcheckDataSourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Namespace          types.String `tfsdk:"namespace"`
	Description        types.String `tfsdk:"description"`
	HealthCheckType    types.String `tfsdk:"healthcheck_type"`
	Timeout            types.Int64  `tfsdk:"timeout"`
	Interval           types.Int64  `tfsdk:"interval"`
	UnhealthyThreshold types.Int64  `tfsdk:"unhealthy_threshold"`
	HealthyThreshold   types.Int64  `tfsdk:"healthy_threshold"`
}

type APIHealthcheck struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		HTTPHealthCheck    interface{} `json:"http_health_check,omitempty"`
		TCPHealthCheck     interface{} `json:"tcp_health_check,omitempty"`
		Timeout            int         `json:"timeout,omitempty"`
		Interval           int         `json:"interval,omitempty"`
		UnhealthyThreshold int         `json:"unhealthy_threshold,omitempty"`
		HealthyThreshold   int         `json:"healthy_threshold,omitempty"`
	} `json:"spec"`
}

func (d *HealthcheckDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_healthcheck"
}

func (d *HealthcheckDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Healthcheck.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_healthcheck`" + ` data source retrieves information about an existing healthcheck configuration.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_healthcheck" "example" {
  name      = "my-healthcheck"
  namespace = "my-namespace"
}

output "healthcheck_type" {
  value = data.f5distributedcloud_healthcheck.example.healthcheck_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the healthcheck.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the healthcheck to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the healthcheck is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the healthcheck.",
				Computed:    true,
			},
			"healthcheck_type": schema.StringAttribute{
				Description: "Type of healthcheck (HTTP or TCP).",
				Computed:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "Timeout in seconds for each healthcheck request.",
				Computed:    true,
			},
			"interval": schema.Int64Attribute{
				Description: "Interval in seconds between healthcheck requests.",
				Computed:    true,
			},
			"unhealthy_threshold": schema.Int64Attribute{
				Description: "Number of consecutive failed checks before marking unhealthy.",
				Computed:    true,
			},
			"healthy_threshold": schema.Int64Attribute{
				Description: "Number of consecutive successful checks before marking healthy.",
				Computed:    true,
			},
		},
	}
}

func (d *HealthcheckDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *HealthcheckDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HealthcheckDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIHealthcheck
	path := fmt.Sprintf("/config/namespaces/%s/healthchecks/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Healthcheck",
			fmt.Sprintf("Could not read healthcheck %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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

	// Determine healthcheck type
	if apiResp.Spec.HTTPHealthCheck != nil {
		data.HealthCheckType = types.StringValue("HTTP")
	} else if apiResp.Spec.TCPHealthCheck != nil {
		data.HealthCheckType = types.StringValue("TCP")
	} else {
		data.HealthCheckType = types.StringValue("unknown")
	}

	data.Timeout = types.Int64Value(int64(apiResp.Spec.Timeout))
	data.Interval = types.Int64Value(int64(apiResp.Spec.Interval))
	data.UnhealthyThreshold = types.Int64Value(int64(apiResp.Spec.UnhealthyThreshold))
	data.HealthyThreshold = types.Int64Value(int64(apiResp.Spec.HealthyThreshold))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
