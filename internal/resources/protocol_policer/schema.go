// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package protocol_policer

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Protocol Policer.",
		MarkdownDescription: `
The ` + "`f5xc_protocol_policer`" + ` resource manages Protocol Policers in F5 Distributed Cloud.

Protocol Policers define rate limiting policies for specific protocols.

## Example Usage

` + "```hcl" + `
resource "f5xc_protocol_policer" "example" {
  name              = "http-policer"
  namespace         = "my-namespace"
  protocol          = "HTTP"
  requests_per_unit = 1000
  unit              = "SECOND"
  burst_size        = 100
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"protocol": schema.StringAttribute{
				Required:    true,
				Description: "Protocol to police (HTTP, HTTPS, TCP, UDP).",
				Validators:  []validator.String{stringvalidator.OneOf("HTTP", "HTTPS", "TCP", "UDP")},
			},
			"requests_per_unit": schema.Int64Attribute{
				Required:    true,
				Description: "Number of requests allowed per time unit.",
			},
			"unit": schema.StringAttribute{
				Required:    true,
				Description: "Time unit for rate limiting (SECOND, MINUTE).",
				Validators:  []validator.String{stringvalidator.OneOf("SECOND", "MINUTE")},
			},
			"burst_size": schema.Int64Attribute{
				Optional:    true,
				Description: "Burst size above the rate limit.",
			},
		},
	}
}
