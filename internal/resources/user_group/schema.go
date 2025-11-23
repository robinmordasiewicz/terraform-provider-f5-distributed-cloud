// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package user_group

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the User Group resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud User Group. User Groups organize users for access control.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_user_group`" + ` resource manages User Groups in F5 Distributed Cloud.

User Groups organize users for access control, allowing you to assign permissions
and policies to groups of users rather than individual users.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_user_group" "example" {
  name        = "my-user-group"
  namespace   = "system"
  description = "User group for administrators"
}
` + "```" + `

## Import

User Groups can be imported using namespace/name:

` + "```shell" + `
terraform import f5_distributed_cloud_user_group.example system/my-user-group
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the User Group.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the User Group.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the User Group is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the User Group.",
				Optional:    true,
			},
		},
	}
}
