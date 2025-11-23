// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package rbac_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud RBAC Policy. RBAC Policies define role-based access control rules.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_rbac_policy`" + ` resource manages RBAC Policies in F5 Distributed Cloud.

RBAC Policies define role-based access control rules that determine what actions users and services can perform on specific resources.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_rbac_policy" "example" {
  name        = "my-rbac-policy"
  namespace   = "my-namespace"
  description = "RBAC policy for application access"

  rules {
    action     = "allow"
    api_groups = ["app.f5.com"]
    resources  = ["virtual_hosts", "origin_pools"]
    verbs      = ["get", "list", "create", "update", "delete"]
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
			"rules": schema.ListNestedBlock{
				Description: "List of RBAC rules defining access permissions.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"action":     schema.StringAttribute{Required: true, Description: "The action to take (allow or deny)."},
						"api_groups": schema.ListAttribute{Optional: true, ElementType: types.StringType, Description: "List of API groups this rule applies to."},
						"resources":  schema.ListAttribute{Optional: true, ElementType: types.StringType, Description: "List of resources this rule applies to."},
						"verbs":      schema.ListAttribute{Optional: true, ElementType: types.StringType, Description: "List of verbs (actions) allowed on the resources."},
					},
				},
			},
		},
	}
}
