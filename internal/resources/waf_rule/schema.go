// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package waf_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud WAF Rule.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_waf_rule`" + ` resource manages WAF Rules in F5 Distributed Cloud.

WAF Rules define custom Web Application Firewall rules for security policies.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_waf_rule" "example" {
  name      = "custom-waf-rule"
  namespace = "my-namespace"
  mode      = "BLOCKING"
  rule_id   = "942100"
  action    = "BLOCK"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"mode": schema.StringAttribute{
				Required:    true,
				Description: "WAF rule mode (BLOCKING, MONITORING).",
				Validators:  []validator.String{stringvalidator.OneOf("BLOCKING", "MONITORING")},
			},
			"rule_id": schema.StringAttribute{
				Optional:    true,
				Description: "WAF rule identifier.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to take (BLOCK, ALLOW, LOG).",
				Validators:  []validator.String{stringvalidator.OneOf("BLOCK", "ALLOW", "LOG")},
			},
		},
	}
}
