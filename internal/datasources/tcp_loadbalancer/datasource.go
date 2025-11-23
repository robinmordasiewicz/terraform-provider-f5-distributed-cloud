// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tcp_loadbalancer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &TCPLoadBalancerDataSource{}

func NewTCPLoadBalancerDataSource() datasource.DataSource {
	return &TCPLoadBalancerDataSource{}
}

type TCPLoadBalancerDataSource struct {
	client *client.Client
}

type TCPLoadBalancerDataSourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Namespace          types.String `tfsdk:"namespace"`
	Description        types.String `tfsdk:"description"`
	ListenPort         types.Int64  `tfsdk:"listen_port"`
	DNSVolterraManaged types.Bool   `tfsdk:"dns_volterra_managed"`
}

type APITCPLoadBalancer struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		ListenPort         int  `json:"listen_port,omitempty"`
		DNSVolterraManaged bool `json:"dns_volterra_managed,omitempty"`
	} `json:"spec"`
}

func (d *TCPLoadBalancerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tcp_loadbalancer"
}

func (d *TCPLoadBalancerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud TCP Load Balancer.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_tcp_loadbalancer`" + ` data source retrieves information about an existing TCP load balancer.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_tcp_loadbalancer" "example" {
  name      = "my-tcp-lb"
  namespace = "my-namespace"
}

output "tcp_lb_port" {
  value = data.f5distributedcloud_tcp_loadbalancer.example.listen_port
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the TCP load balancer.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the TCP load balancer to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the TCP load balancer is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the TCP load balancer.",
				Computed:    true,
			},
			"listen_port": schema.Int64Attribute{
				Description: "Port on which the TCP load balancer listens.",
				Computed:    true,
			},
			"dns_volterra_managed": schema.BoolAttribute{
				Description: "Whether DNS is managed by Volterra.",
				Computed:    true,
			},
		},
	}
}

func (d *TCPLoadBalancerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TCPLoadBalancerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TCPLoadBalancerDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APITCPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/tcp_loadbalancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading TCP Load Balancer",
			fmt.Sprintf("Could not read TCP load balancer %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	data.ListenPort = types.Int64Value(int64(apiResp.Spec.ListenPort))
	data.DNSVolterraManaged = types.BoolValue(apiResp.Spec.DNSVolterraManaged)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
