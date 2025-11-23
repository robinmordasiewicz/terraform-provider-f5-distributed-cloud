// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package malicious_user_mitigation

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Malicious User Mitigation.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_malicious_user_mitigation`" + ` resource manages Malicious User Mitigation policies in F5 Distributed Cloud.

Malicious User Mitigation helps protect against malicious user behavior through various challenge and blocking mechanisms.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_malicious_user_mitigation" "example" {
  name                 = "malicious-user-policy"
  namespace            = "my-namespace"
  mitigation_type      = "CHALLENGE"
  captcha_challenge    = true
  javascript_challenge = true
  temporary_block      = false
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"mitigation_type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of mitigation (CHALLENGE, BLOCK).",
			},
			"captcha_challenge": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable CAPTCHA challenge for suspicious users.",
			},
			"javascript_challenge": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable JavaScript challenge for suspicious users.",
			},
			"temporary_block": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable temporary blocking of malicious users.",
			},
			"block_duration": schema.Int64Attribute{
				Optional:    true,
				Description: "Duration in seconds to block malicious users.",
			},
		},
	}
}
