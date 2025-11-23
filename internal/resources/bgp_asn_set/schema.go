// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package bgp_asn_set

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud BGP ASN Set.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_bgp_asn_set`" + ` resource manages BGP ASN Sets in F5 Distributed Cloud.

BGP ASN Sets define collections of Autonomous System Numbers for BGP routing policies.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_bgp_asn_set" "example" {
  name        = "my-asn-set"
  namespace   = "system"
  description = "BGP ASN set for routing"
  asns        = [65000, 65001, 65002]
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"asns":        schema.ListAttribute{Optional: true, ElementType: types.Int64Type, Description: "List of Autonomous System Numbers."},
		},
	}
}
