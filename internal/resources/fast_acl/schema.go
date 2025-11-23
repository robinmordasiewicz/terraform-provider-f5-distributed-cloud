// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package fast_acl

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Fast ACL.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_fast_acl`" + ` resource manages Fast ACLs in F5 Distributed Cloud.

Fast ACLs provide high-performance access control lists.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_fast_acl" "example" {
  name        = "my-fast-acl"
  namespace   = "system"
  description = "High-performance access control list"
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
