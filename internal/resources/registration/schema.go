// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package registration

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Registration resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Registration. Registrations manage site registration processes.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_registration`" + ` resource manages Registrations in F5 Distributed Cloud.

Registrations manage site registration processes, allowing sites to be registered
and configured within the F5 Distributed Cloud platform.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_registration" "example" {
  name        = "my-registration"
  namespace   = "system"
  description = "My site registration"
}
` + "```" + `

## Import

Registrations can be imported using namespace/name:

` + "```shell" + `
terraform import f5_distributed_cloud_registration.example system/my-registration
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Registration.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Registration.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Registration is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Registration.",
				Optional:    true,
			},
		},
	}
}
