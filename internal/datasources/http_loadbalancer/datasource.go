// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package http_loadbalancer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &HTTPLoadBalancerDataSource{}

func NewHTTPLoadBalancerDataSource() datasource.DataSource {
	return &HTTPLoadBalancerDataSource{}
}

type HTTPLoadBalancerDataSource struct {
	client *client.Client
}

type HTTPLoadBalancerDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Domains     types.List   `tfsdk:"domains"`
}

type APIHTTPLoadBalancer struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace,omitempty"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Domains []string `json:"domains,omitempty"`
	} `json:"spec"`
}

func (d *HTTPLoadBalancerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_http_loadbalancer"
}

func (d *HTTPLoadBalancerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud HTTP Load Balancer.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_http_loadbalancer`" + ` data source retrieves information about an existing HTTP Load Balancer.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_http_loadbalancer" "example" {
  name      = "my-lb"
  namespace = "my-namespace"
}

output "lb_domains" {
  value = data.f5distributedcloud_http_loadbalancer.example.domains
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the HTTP Load Balancer.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the HTTP Load Balancer to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace of the HTTP Load Balancer.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the HTTP Load Balancer.",
				Computed:    true,
			},
			"domains": schema.ListAttribute{
				Description: "List of domains configured on the HTTP Load Balancer.",
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *HTTPLoadBalancerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *HTTPLoadBalancerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HTTPLoadBalancerDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIHTTPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/http_loadbalancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading HTTP Load Balancer",
			fmt.Sprintf("Could not read HTTP Load Balancer %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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

	// Convert domains
	if len(apiResp.Spec.Domains) > 0 {
		domainsList, diags := types.ListValueFrom(ctx, types.StringType, apiResp.Spec.Domains)
		resp.Diagnostics.Append(diags...)
		data.Domains = domainsList
	} else {
		data.Domains = types.ListNull(types.StringType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
