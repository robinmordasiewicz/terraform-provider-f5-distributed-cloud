// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package healthcheck

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Schema returns the schema for the Healthcheck resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Health Check.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_healthcheck`" + ` resource manages Health Checks in F5 Distributed Cloud.

Health Checks define how origin servers are monitored to determine their health status.
They can be referenced by origin pools to enable automatic traffic routing away from unhealthy servers.

## Example Usage

### HTTP Health Check

` + "```hcl" + `
resource "f5distributedcloud_healthcheck" "http" {
  name        = "my-http-healthcheck"
  namespace   = "my-namespace"
  description = "HTTP health check"

  http_health_check {
    path      = "/health"
    use_https = false
  }

  timeout             = 3
  interval            = 15
  unhealthy_threshold = 3
  healthy_threshold   = 2
}
` + "```" + `

### TCP Health Check

` + "```hcl" + `
resource "f5distributedcloud_healthcheck" "tcp" {
  name        = "my-tcp-healthcheck"
  namespace   = "my-namespace"
  description = "TCP health check"

  tcp_health_check {}

  timeout             = 3
  interval            = 15
  unhealthy_threshold = 3
  healthy_threshold   = 2
}
` + "```" + `

## Import

Health Checks can be imported using namespace/name:

` + "```shell" + `
terraform import f5distributedcloud_healthcheck.example my-namespace/my-healthcheck
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Health Check.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Health Check.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Health Check is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Health Check.",
				Optional:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "Timeout for health check in seconds. Default is 3.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3),
			},
			"interval": schema.Int64Attribute{
				Description: "Interval between health checks in seconds. Default is 15.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(15),
			},
			"unhealthy_threshold": schema.Int64Attribute{
				Description: "Number of consecutive failures before marking unhealthy. Default is 3.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3),
			},
			"healthy_threshold": schema.Int64Attribute{
				Description: "Number of consecutive successes before marking healthy. Default is 2.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(2),
			},
		},
		Blocks: map[string]schema.Block{
			"http_health_check": schema.SingleNestedBlock{
				Description: "HTTP health check configuration.",
				Attributes: map[string]schema.Attribute{
					"path": schema.StringAttribute{
						Description: "HTTP path to check. Default is '/'.",
						Optional:    true,
					},
					"use_https": schema.BoolAttribute{
						Description: "Use HTTPS for health check.",
						Optional:    true,
					},
					"expected_status_codes": schema.ListAttribute{
						Description: "List of expected HTTP status codes (e.g., '200', '2xx').",
						Optional:    true,
						ElementType: types.StringType,
					},
					"headers": schema.MapAttribute{
						Description: "HTTP headers to send with health check request.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"tcp_health_check": schema.SingleNestedBlock{
				Description: "TCP health check configuration (connectivity check).",
				Attributes:  map[string]schema.Attribute{},
			},
		},
	}
}
