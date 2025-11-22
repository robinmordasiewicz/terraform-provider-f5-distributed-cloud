// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package ip_prefix_set

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud IP Prefix Set.",
		MarkdownDescription: `
The ` + "`f5xc_ip_prefix_set`" + ` resource manages IP Prefix Sets in F5 Distributed Cloud.

IP Prefix Sets define a collection of IP address prefixes for use in security rules and policies.

## Example Usage

` + "```hcl" + `
resource "f5xc_ip_prefix_set" "example" {
  name      = "trusted-networks"
  namespace = "my-namespace"
  prefixes  = ["10.0.0.0/8", "192.168.0.0/16", "172.16.0.0/12"]
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"prefixes": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "List of IP prefixes in CIDR notation.",
			},
		},
	}
}
