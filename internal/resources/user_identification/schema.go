// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package user_identification

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the User Identification resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud User Identification. User Identifications define how users are identified.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_user_identification`" + ` resource manages User Identifications in F5 Distributed Cloud.

User Identifications define how users are identified for tracking and security purposes.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_user_identification" "example" {
  name        = "my-user-identification"
  namespace   = "my-namespace"
  description = "Example User Identification"
}
` + "```" + `

## Import

User Identifications can be imported using namespace/name:

` + "```shell" + `
terraform import f5_distributed_cloud_user_identification.example my-namespace/my-user-identification
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the User Identification.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the User Identification.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the User Identification is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the User Identification.",
				Optional:    true,
			},
		},
	}
}
