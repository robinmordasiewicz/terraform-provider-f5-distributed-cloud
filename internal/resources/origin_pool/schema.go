// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package origin_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Schema returns the schema for the Origin Pool resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Origin Pool.",
		MarkdownDescription: `
The ` + "`f5xc_origin_pool`" + ` resource manages Origin Pools in F5 Distributed Cloud.

Origin Pools define groups of backend servers that serve application traffic.
They are used by HTTP Load Balancers to route traffic to appropriate backends.

## Example Usage

` + "```hcl" + `
resource "f5xc_origin_pool" "example" {
  name        = "my-origin-pool"
  namespace   = "my-namespace"
  description = "Example Origin Pool"

  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }
}
` + "```" + `

## Import

Origin Pools can be imported using namespace/name:

` + "```shell" + `
terraform import f5xc_origin_pool.example my-namespace/my-origin-pool
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Origin Pool.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Origin Pool.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Origin Pool is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Origin Pool.",
				Optional:    true,
			},
			"port": schema.Int64Attribute{
				Description: "Port on which origin servers listen.",
				Required:    true,
			},
			"endpoint_protocol": schema.StringAttribute{
				Description: "Protocol to use for connecting to origin servers: 'http' or 'https'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("http"),
				Validators: []validator.String{
					stringvalidator.OneOf("http", "https"),
				},
			},
			"health_check_port": schema.Int64Attribute{
				Description: "Port for health checks. Defaults to the origin port if not specified.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"loadbalancer_algorithm": schema.StringAttribute{
				Description: "Load balancing algorithm: 'ROUND_ROBIN', 'LEAST_REQUEST', 'RANDOM', 'SOURCE_IP_HASH'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ROUND_ROBIN"),
				Validators: []validator.String{
					stringvalidator.OneOf("ROUND_ROBIN", "LEAST_REQUEST", "RANDOM", "SOURCE_IP_HASH"),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"origin_servers": schema.ListNestedBlock{
				Description: "List of origin servers in the pool.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"labels": schema.MapAttribute{
							Description: "Labels for the origin server.",
							Optional:    true,
							ElementType: types.StringType,
						},
					},
					Blocks: map[string]schema.Block{
						"public_ip": schema.SingleNestedBlock{
							Description: "Public IP address of the origin server.",
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "IPv4 or IPv6 address.",
									Required:    true,
								},
							},
						},
						"private_ip": schema.SingleNestedBlock{
							Description: "Private IP address of the origin server.",
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description: "IPv4 or IPv6 address.",
									Required:    true,
								},
								"site_name": schema.StringAttribute{
									Description: "Name of the site where the server is located.",
									Required:    true,
								},
								"inside_network": schema.BoolAttribute{
									Description: "Whether to use inside network.",
									Optional:    true,
								},
							},
						},
						"public_name": schema.SingleNestedBlock{
							Description: "DNS name of the origin server.",
							Attributes: map[string]schema.Attribute{
								"dns_name": schema.StringAttribute{
									Description: "Fully qualified domain name.",
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}
