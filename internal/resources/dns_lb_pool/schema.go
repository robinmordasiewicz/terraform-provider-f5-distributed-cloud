// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_lb_pool

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud DNS Load Balancer Pool.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_dns_lb_pool`" + ` resource manages DNS Load Balancer Pools in F5 Distributed Cloud.

DNS LB Pools define groups of endpoints for DNS-based load balancing with health checking and failover capabilities.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_dns_lb_pool" "example" {
  name      = "my-dns-pool"
  namespace = "my-namespace"

  load_balancing_mode = "round-robin"

  members {
    address  = "10.0.0.1"
    port     = 80
    priority = 1
    ratio    = 1
    enabled  = true
  }

  members {
    address  = "10.0.0.2"
    port     = 80
    priority = 1
    ratio    = 1
    enabled  = true
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"load_balancing_mode": schema.StringAttribute{
				Optional:    true,
				Description: "Load balancing mode for the pool.",
				Validators:  []validator.String{stringvalidator.OneOf("round-robin", "ratio-member", "least-connections", "priority")},
			},
		},
		Blocks: map[string]schema.Block{
			"members": schema.ListNestedBlock{
				Description: "Pool members configuration.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"address":  schema.StringAttribute{Required: true, Description: "IP address of the pool member."},
						"port":     schema.Int64Attribute{Required: true, Description: "Port number of the pool member."},
						"priority": schema.Int64Attribute{Optional: true, Description: "Priority of the pool member."},
						"ratio":    schema.Int64Attribute{Optional: true, Description: "Ratio weight for load balancing."},
						"enabled":  schema.BoolAttribute{Optional: true, Description: "Whether the pool member is enabled."},
					},
				},
			},
		},
	}
}
