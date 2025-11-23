// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package origin_pool

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &OriginPoolDataSource{}

func NewOriginPoolDataSource() datasource.DataSource {
	return &OriginPoolDataSource{}
}

type OriginPoolDataSource struct {
	client *client.Client
}

type OriginPoolDataSourceModel struct {
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Namespace             types.String `tfsdk:"namespace"`
	Description           types.String `tfsdk:"description"`
	Port                  types.Int64  `tfsdk:"port"`
	EndpointProtocol      types.String `tfsdk:"endpoint_protocol"`
	LoadbalancerAlgorithm types.String `tfsdk:"loadbalancer_algorithm"`
	HealthCheckPort       types.Int64  `tfsdk:"health_check_port"`
}

type APIOriginPool struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace,omitempty"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Port                  int64  `json:"port,omitempty"`
		EndpointProtocol      string `json:"endpoint_protocol,omitempty"`
		LoadbalancerAlgorithm string `json:"loadbalancer_algorithm,omitempty"`
		HealthCheckPort       int64  `json:"health_check_port,omitempty"`
	} `json:"spec"`
}

func (d *OriginPoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_origin_pool"
}

func (d *OriginPoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Origin Pool.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_origin_pool`" + ` data source retrieves information about an existing origin pool.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_origin_pool" "example" {
  name      = "my-origin-pool"
  namespace = "my-namespace"
}

output "origin_pool_port" {
  value = data.f5distributedcloud_origin_pool.example.port
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the origin pool.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the origin pool to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace containing the origin pool.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the origin pool.",
				Computed:    true,
			},
			"port": schema.Int64Attribute{
				Description: "Port on which origin servers listen.",
				Computed:    true,
			},
			"endpoint_protocol": schema.StringAttribute{
				Description: "Protocol to use for connecting to origin servers.",
				Computed:    true,
			},
			"loadbalancer_algorithm": schema.StringAttribute{
				Description: "Load balancing algorithm used.",
				Computed:    true,
			},
			"health_check_port": schema.Int64Attribute{
				Description: "Port for health checks.",
				Computed:    true,
			},
		},
	}
}

func (d *OriginPoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OriginPoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OriginPoolDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIOriginPool
	path := fmt.Sprintf("/config/namespaces/%s/origin_pools/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Origin Pool",
			fmt.Sprintf("Could not read origin pool %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err))
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

	data.Port = types.Int64Value(apiResp.Spec.Port)

	if apiResp.Spec.EndpointProtocol != "" {
		data.EndpointProtocol = types.StringValue(apiResp.Spec.EndpointProtocol)
	}

	if apiResp.Spec.LoadbalancerAlgorithm != "" {
		data.LoadbalancerAlgorithm = types.StringValue(apiResp.Spec.LoadbalancerAlgorithm)
	}

	if apiResp.Spec.HealthCheckPort > 0 {
		data.HealthCheckPort = types.Int64Value(apiResp.Spec.HealthCheckPort)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
