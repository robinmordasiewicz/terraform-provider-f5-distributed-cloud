// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud User.",
		MarkdownDescription: `
The ` + "`f5xc_user`" + ` resource manages Users in F5 Distributed Cloud.

Users represent individuals who can access the F5 Distributed Cloud console and API.

## Example Usage

` + "```hcl" + `
resource "f5xc_user" "example" {
  name       = "john-doe"
  namespace  = "system"
  email      = "john.doe@example.com"
  first_name = "John"
  last_name  = "Doe"
  type       = "USER"

  roles {
    name      = "admin-role"
    namespace = "system"
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"email":       schema.StringAttribute{Required: true, Description: "Email address of the user."},
			"first_name":  schema.StringAttribute{Optional: true, Description: "First name of the user."},
			"last_name":   schema.StringAttribute{Optional: true, Description: "Last name of the user."},
			"type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of user (USER, SERVICE_CREDENTIAL).",
				Validators:  []validator.String{stringvalidator.OneOf("USER", "SERVICE_CREDENTIAL")},
			},
		},
		Blocks: map[string]schema.Block{
			"roles": schema.ListNestedBlock{
				Description: "List of roles assigned to the user.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name":      schema.StringAttribute{Required: true, Description: "Name of the role."},
						"namespace": schema.StringAttribute{Required: true, Description: "Namespace of the role."},
					},
				},
			},
		},
	}
}
