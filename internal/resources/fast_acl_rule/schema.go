// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package fast_acl_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Fast ACL Rule.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_fast_acl_rule`" + ` resource manages Fast ACL Rules in F5 Distributed Cloud.

Fast ACL Rules define individual access control entries.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "my-namespace"
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
