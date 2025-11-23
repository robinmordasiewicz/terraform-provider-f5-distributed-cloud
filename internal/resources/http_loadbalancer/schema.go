// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package http_loadbalancer

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Schema returns the schema for the HTTP Load Balancer resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud HTTP Load Balancer.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_http_loadbalancer`" + ` resource manages HTTP Load Balancers in F5 Distributed Cloud.

HTTP Load Balancers distribute incoming traffic across multiple origin servers, providing
high availability, scalability, and improved performance for web applications.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_http_loadbalancer" "example" {
  name        = "my-http-lb"
  namespace   = "my-namespace"
  description = "Example HTTP Load Balancer"

  domains = ["app.example.com"]

  http_type = "http"
  advertise_port = 80

  default_origin_pools {
    pool_name      = "my-origin-pool"
    pool_namespace = "my-namespace"
    weight         = 1
    priority       = 1
  }
}
` + "```" + `

## Import

HTTP Load Balancers can be imported using namespace/name:

` + "```shell" + `
terraform import f5distributedcloud_http_loadbalancer.example my-namespace/my-http-lb
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the HTTP Load Balancer.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the HTTP Load Balancer.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the HTTP Load Balancer is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the HTTP Load Balancer.",
				Optional:    true,
			},
			"domains": schema.ListAttribute{
				Description: "List of domains for the HTTP Load Balancer.",
				Required:    true,
				ElementType: schema.StringAttribute{}.GetType(),
			},
			"http_type": schema.StringAttribute{
				Description: "Type of HTTP Load Balancer: 'http' or 'https'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("http"),
				Validators: []validator.String{
					stringvalidator.OneOf("http", "https"),
				},
			},
			"advertise_port": schema.Int64Attribute{
				Description: "Port to advertise the HTTP Load Balancer on.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(80),
			},
			"disable_waf": schema.BoolAttribute{
				Description: "Whether to disable WAF for this load balancer.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
		Blocks: map[string]schema.Block{
			"default_origin_pools": schema.ListNestedBlock{
				Description: "Default origin pools for the HTTP Load Balancer.",
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
							Description: "Weight for load balancing.",
							Optional:    true,
							Computed:    true,
							Default:     int64default.StaticInt64(1),
						},
						"priority": schema.Int64Attribute{
							Description: "Priority for the origin pool.",
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
