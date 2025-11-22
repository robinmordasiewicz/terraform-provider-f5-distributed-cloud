// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_lb_pool

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &DNSLBPoolDataSource{}

func NewDNSLBPoolDataSource() datasource.DataSource {
	return &DNSLBPoolDataSource{}
}

type DNSLBPoolDataSource struct {
	client *client.Client
}

type DNSLBPoolDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	PoolType    types.String `tfsdk:"pool_type"`
}

type APIDNSLBPool struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		PoolType string `json:"pool_type,omitempty"`
	} `json:"spec"`
}

func (d *DNSLBPoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_lb_pool"
}

func (d *DNSLBPoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud DNS Load Balancer Pool.",
		MarkdownDescription: `
The ` + "`f5xc_dns_lb_pool`" + ` data source retrieves information about an existing DNS load balancer pool.

## Example Usage

` + "```hcl" + `
data "f5xc_dns_lb_pool" "example" {
  name      = "my-dns-lb-pool"
  namespace = "my-namespace"
}

output "pool_type" {
  value = data.f5xc_dns_lb_pool.example.pool_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the DNS LB pool.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the DNS LB pool to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the DNS LB pool is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the DNS LB pool.",
				Computed:    true,
			},
			"pool_type": schema.StringAttribute{
				Description: "Type of the DNS LB pool.",
				Computed:    true,
			},
		},
	}
}

func (d *DNSLBPoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DNSLBPoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DNSLBPoolDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIDNSLBPool
	path := fmt.Sprintf("/config/namespaces/%s/dns_lb_pools/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading DNS LB Pool",
			fmt.Sprintf("Could not read DNS LB pool %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.PoolType != "" {
		data.PoolType = types.StringValue(apiResp.Spec.PoolType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
