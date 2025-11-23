// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package active_service_policies

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Active Service Policies.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_active_service_policies`" + ` resource manages Active Service Policies in F5 Distributed Cloud.

Active Service Policies allow you to activate and apply a set of service policies to sites or virtual sites.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_active_service_policies" "example" {
  name      = "active-policies"
  namespace = "system"

  policies {
    name      = "rate-limit-policy"
    namespace = "my-namespace"
  }

  policies {
    name      = "waf-policy"
    namespace = "my-namespace"
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
		},
		Blocks: map[string]schema.Block{
			"policies": schema.ListNestedBlock{
				Description: "List of service policies to activate.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name":      schema.StringAttribute{Required: true, Description: "Name of the service policy."},
						"namespace": schema.StringAttribute{Required: true, Description: "Namespace of the service policy."},
					},
				},
			},
		},
	}
}
