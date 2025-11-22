// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package nat_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud NAT Policy.",
		MarkdownDescription: `
The ` + "`f5xc_nat_policy`" + ` resource manages NAT Policies in F5 Distributed Cloud.

NAT Policies configure network address translation rules.

## Example Usage

` + "```hcl" + `
resource "f5xc_nat_policy" "example" {
  name        = "nat-policy"
  namespace   = "my-namespace"
  description = "Example NAT policy"
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
