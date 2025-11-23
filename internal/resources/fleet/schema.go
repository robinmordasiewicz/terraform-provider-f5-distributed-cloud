// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package fleet

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Fleet. Fleets group and manage edge sites.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_fleet`" + ` resource manages Fleets in F5 Distributed Cloud.

Fleets group and manage edge sites.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_fleet" "example" {
  name        = "my-fleet"
  namespace   = "my-namespace"
  description = "Example fleet"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
		},
	}
}
