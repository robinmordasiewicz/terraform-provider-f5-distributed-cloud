// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client_connection_limit

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Client Connection Limit.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_client_connection_limit`" + ` resource manages Client Connection Limits in F5 Distributed Cloud.

Client Connection Limits define connection restrictions and circuit breaker settings.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_client_connection_limit" "example" {
  name                 = "api-conn-limit"
  namespace            = "my-namespace"
  max_connections      = 10000
  connection_timeout   = 60
  enable_circuit_break = true
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"max_connections": schema.Int64Attribute{
				Required:    true,
				Description: "Maximum number of client connections allowed.",
			},
			"connection_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "Connection timeout in seconds.",
			},
			"enable_circuit_break": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable circuit breaker functionality.",
			},
		},
	}
}
