// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_lb_health_check

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud DNS Load Balancer Health Check.",
		MarkdownDescription: `
The ` + "`f5xc_dns_lb_health_check`" + ` resource manages DNS Load Balancer Health Checks in F5 Distributed Cloud.

Health checks monitor the availability of DNS load balancer endpoints.

## Example Usage

` + "```hcl" + `
resource "f5xc_dns_lb_health_check" "example" {
  name        = "my-health-check"
  namespace   = "system"
  description = "HTTP health check"
  protocol    = "HTTP"
  port        = 80
  interval    = 30
  timeout     = 10
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"protocol":    schema.StringAttribute{Required: true, Description: "Health check protocol (HTTP, HTTPS, TCP)."},
			"port":        schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(80), Description: "Port to check."},
			"interval":    schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(30), Description: "Interval between health checks in seconds."},
			"timeout":     schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(10), Description: "Health check timeout in seconds."},
		},
	}
}
