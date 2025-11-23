// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package subnet

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Subnet.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_subnet`" + ` resource manages Subnets in F5 Distributed Cloud.

Subnets define IP address ranges within virtual networks.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_subnet" "example" {
  name        = "my-subnet"
  namespace   = "system"
  description = "Application subnet"
  cidr        = "10.0.1.0/24"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"cidr":        schema.StringAttribute{Required: true, Description: "CIDR block for the subnet."},
		},
	}
}
