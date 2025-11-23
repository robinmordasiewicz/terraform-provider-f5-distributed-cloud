// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package route

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Route.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_route`" + ` resource manages Routes in F5 Distributed Cloud.

Routes define traffic routing rules for load balancers.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_route" "example" {
  name        = "api-route"
  namespace   = "my-namespace"
  path        = "/api/v1/*"
  methods     = ["GET", "POST"]
  destination = "backend-pool"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"path": schema.StringAttribute{
				Required:    true,
				Description: "URL path pattern for the route.",
			},
			"methods": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "HTTP methods allowed for this route.",
			},
			"destination": schema.StringAttribute{
				Required:    true,
				Description: "Destination pool or service for matched traffic.",
			},
		},
	}
}
