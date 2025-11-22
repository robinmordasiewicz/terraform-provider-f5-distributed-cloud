// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package token

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Token resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Token. Tokens provide authentication for site registration.",
		MarkdownDescription: `
The ` + "`f5xc_token`" + ` resource manages Tokens in F5 Distributed Cloud.

Tokens provide authentication for site registration, allowing Customer Edge (CE)
sites to securely register with the F5 Distributed Cloud platform.

## Example Usage

` + "```hcl" + `
resource "f5xc_token" "example" {
  name        = "my-site-token"
  namespace   = "system"
  description = "Token for site registration"
}
` + "```" + `

## Import

Tokens can be imported using namespace/name:

` + "```shell" + `
terraform import f5xc_token.example system/my-site-token
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Token.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Token.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Token is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Token.",
				Optional:    true,
			},
		},
	}
}
