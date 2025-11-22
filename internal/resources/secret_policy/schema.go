// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package secret_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Secret Policy.",
		MarkdownDescription: `
The ` + "`f5xc_secret_policy`" + ` resource manages Secret Policies in F5 Distributed Cloud.

Secret Policies define access control for secrets management.

## Example Usage

` + "```hcl" + `
resource "f5xc_secret_policy" "example" {
  name        = "app-secrets-policy"
  namespace   = "my-namespace"
  policy_type = "ALLOW_READ"
  rules       = "namespace/*"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"policy_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of secret policy (ALLOW_READ, ALLOW_WRITE, DENY).",
				Validators:  []validator.String{stringvalidator.OneOf("ALLOW_READ", "ALLOW_WRITE", "DENY")},
			},
			"rules": schema.StringAttribute{
				Optional:    true,
				Description: "Policy rules pattern.",
			},
		},
	}
}
