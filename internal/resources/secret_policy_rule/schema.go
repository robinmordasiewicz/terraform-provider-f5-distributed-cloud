// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package secret_policy_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Secret Policy Rule.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_secret_policy_rule`" + ` resource manages Secret Policy Rules in F5 Distributed Cloud.

Secret Policy Rules define specific rules within a secret policy.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_secret_policy_rule" "example" {
  name        = "my-secret-policy-rule"
  namespace   = "my-namespace"
  description = "Example secret policy rule"
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
