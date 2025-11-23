// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package aws_tgw_site

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud AWS TGW Site.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_aws_tgw_site`" + ` resource manages AWS Transit Gateway Sites in F5 Distributed Cloud.

AWS TGW Sites enable connectivity between F5 XC and AWS Transit Gateway infrastructure.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_aws_tgw_site" "example" {
  name        = "my-aws-tgw-site"
  namespace   = "system"
  description = "Example AWS TGW site"
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
