// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package bgp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud BGP configuration.",
		MarkdownDescription: `
The ` + "`f5xc_bgp`" + ` resource manages BGP configurations in F5 Distributed Cloud.

BGP configurations define Border Gateway Protocol settings for network routing.

## Example Usage

` + "```hcl" + `
resource "f5xc_bgp" "example" {
  name        = "my-bgp-config"
  namespace   = "system"
  description = "BGP configuration"
  asn         = 65000
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"asn":         schema.Int64Attribute{Optional: true, Description: "Autonomous System Number for BGP."},
		},
	}
}
