// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace_role

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
		Description: "Manages an F5 Distributed Cloud Namespace Role. Namespace Roles define permissions within a namespace.",
		MarkdownDescription: `
The ` + "`f5xc_namespace_role`" + ` resource manages Namespace Roles in F5 Distributed Cloud.

Namespace Roles define permissions within a namespace, allowing fine-grained access control for resources and operations.

## Example Usage

` + "```hcl" + `
resource "f5xc_namespace_role" "viewer" {
  name        = "viewer-role"
  namespace   = "my-namespace"
  description = "Read-only access role for namespace resources"
  role_type   = "CUSTOM"

  permissions {
    resource   = "origin_pool"
    operations = ["read", "list"]
  }

  permissions {
    resource   = "http_loadbalancer"
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
				Description: "Type of namespace role (CUSTOM, PREDEFINED).",
				Validators:  []validator.String{stringvalidator.OneOf("CUSTOM", "PREDEFINED")},
			},
		},
		Blocks: map[string]schema.Block{
			"permissions": schema.ListNestedBlock{
				Description: "Permission definitions for the namespace role.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"resource": schema.StringAttribute{
							Required:    true,
							Description: "Resource type to grant permissions on within the namespace.",
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
