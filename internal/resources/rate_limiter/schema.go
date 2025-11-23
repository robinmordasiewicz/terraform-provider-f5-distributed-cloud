// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package rate_limiter

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Rate Limiter.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_rate_limiter`" + ` resource manages Rate Limiters in F5 Distributed Cloud.

Rate Limiters define rate limiting policies for controlling request throughput.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_rate_limiter" "example" {
  name             = "api-rate-limit"
  namespace        = "my-namespace"
  total_number     = 100
  unit             = "SECOND"
  burst_multiplier = 2
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"total_number": schema.Int64Attribute{
				Required:    true,
				Description: "Total number of requests allowed in the specified time unit.",
			},
			"unit": schema.StringAttribute{
				Required:    true,
				Description: "Time unit for rate limiting (SECOND, MINUTE).",
				Validators:  []validator.String{stringvalidator.OneOf("SECOND", "MINUTE")},
			},
			"burst_multiplier": schema.Int64Attribute{
				Optional:    true,
				Description: "Multiplier for burst capacity above the rate limit.",
			},
		},
	}
}
