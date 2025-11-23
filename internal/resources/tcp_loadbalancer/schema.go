// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tcp_loadbalancer

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Schema returns the schema for the TCP Load Balancer resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud TCP Load Balancer.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_tcp_loadbalancer`" + ` resource manages TCP Load Balancers in F5 Distributed Cloud.

TCP Load Balancers distribute incoming TCP traffic across origin pools based on
configured rules and policies. They operate at Layer 4 and handle raw TCP connections.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_tcp_loadbalancer" "example" {
  name        = "my-tcp-lb"
  namespace   = "my-namespace"
  description = "Example TCP Load Balancer"

  domains     = ["tcp.example.com"]
  listen_port = 3306

  origin_pools {
    pool_name      = f5_distributed_cloud_origin_pool.mysql.name
    pool_namespace = "my-namespace"
    weight         = 1
    priority       = 1
  }

  advertise_on_public = true
}
` + "```" + `

## Import

TCP Load Balancers can be imported using namespace/name:

` + "```shell" + `
terraform import f5_distributed_cloud_tcp_loadbalancer.example my-namespace/my-tcp-lb
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the TCP Load Balancer.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the TCP Load Balancer.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the TCP Load Balancer is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the TCP Load Balancer.",
				Optional:    true,
			},
			"domains": schema.ListAttribute{
				Description: "List of domains this TCP load balancer serves.",
				Required:    true,
				ElementType: types.StringType,
			},
			"listen_port": schema.Int64Attribute{
				Description: "TCP port to listen on.",
				Required:    true,
			},
			"advertise_on_public": schema.BoolAttribute{
				Description: "Advertise on public network.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"advertise_on_public_default_vip": schema.BoolAttribute{
				Description: "Advertise on public network with default VIP.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"dns_volterra_managed": schema.BoolAttribute{
				Description: "DNS managed by Volterra.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"hash_policy_choice_least_active": schema.BoolAttribute{
				Description: "Use least active hash policy for load balancing.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"hash_policy_choice_random": schema.BoolAttribute{
				Description: "Use random hash policy for load balancing.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"hash_policy_choice_source_ip": schema.BoolAttribute{
				Description: "Use source IP hash policy for load balancing (sticky sessions).",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"idle_timeout": schema.Int64Attribute{
				Description: "Idle timeout in milliseconds. Default is 3600000 (1 hour).",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3600000),
			},
		},
		Blocks: map[string]schema.Block{
			"origin_pools": schema.ListNestedBlock{
				Description: "List of origin pools for this TCP load balancer.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"pool_name": schema.StringAttribute{
							Description: "Name of the origin pool.",
							Required:    true,
						},
						"pool_namespace": schema.StringAttribute{
							Description: "Namespace of the origin pool.",
							Required:    true,
						},
						"weight": schema.Int64Attribute{
							Description: "Weight for load balancing. Default is 1.",
							Optional:    true,
							Computed:    true,
							Default:     int64default.StaticInt64(1),
						},
						"priority": schema.Int64Attribute{
							Description: "Priority for failover. Lower values are higher priority. Default is 1.",
							Optional:    true,
							Computed:    true,
							Default:     int64default.StaticInt64(1),
						},
					},
				},
			},
		},
	}
}
