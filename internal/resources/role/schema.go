// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package role

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Role.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_role`" + ` resource manages Roles in F5 Distributed Cloud.

Roles define permissions for accessing and managing resources within the F5 XC platform.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_role" "viewer" {
  name        = "viewer-role"
  namespace   = "system"
  description = "Read-only access role"
  role_type   = "CUSTOM"

  permissions {
    resource   = "namespace"
    operations = ["read"]
  }

  permissions {
    resource   = "origin_pool"
    operations = ["read", "list"]
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"role_type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of role (CUSTOM, PREDEFINED).",
				Validators:  []validator.String{stringvalidator.OneOf("CUSTOM", "PREDEFINED")},
			},
		},
		Blocks: map[string]schema.Block{
			"permissions": schema.ListNestedBlock{
				Description: "Permission definitions for the role.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"resource": schema.StringAttribute{
							Required:    true,
							Description: "Resource type to grant permissions on.",
						},
						"operations": schema.ListAttribute{
							Required:    true,
							ElementType: types.StringType,
							Description: "Operations allowed (read, write, delete, list).",
						},
					},
				},
			},
		},
	}
}
