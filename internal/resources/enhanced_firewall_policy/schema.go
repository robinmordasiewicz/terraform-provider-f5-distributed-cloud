// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package enhanced_firewall_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Enhanced Firewall Policy.",
		MarkdownDescription: `
The ` + "`f5xc_enhanced_firewall_policy`" + ` resource manages Enhanced Firewall Policies in F5 Distributed Cloud.

Enhanced Firewall Policies provide Layer 3/4 network security controls with advanced rule configurations.

## Example Usage

` + "```hcl" + `
resource "f5xc_enhanced_firewall_policy" "example" {
  name      = "enhanced-fw"
  namespace = "my-namespace"

  rules {
    name     = "allow-https"
    action   = "ALLOW"
    protocol = "TCP"
    dst_port = 443
  }

  rules {
    name   = "deny-all"
    action = "DENY"
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
		},
		Blocks: map[string]schema.Block{
			"rules": schema.ListNestedBlock{
				Description: "Firewall rules configuration.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{Required: true, Description: "Name of the firewall rule."},
						"action": schema.StringAttribute{
							Required:    true,
							Description: "Action to take (ALLOW/DENY).",
							Validators:  []validator.String{stringvalidator.OneOf("ALLOW", "DENY")},
						},
						"protocol": schema.StringAttribute{
							Optional:    true,
							Description: "Protocol to match (TCP/UDP/ICMP/ANY).",
							Validators:  []validator.String{stringvalidator.OneOf("TCP", "UDP", "ICMP", "ANY")},
						},
						"src_ip":   schema.StringAttribute{Optional: true, Description: "Source IP address or CIDR."},
						"dst_ip":   schema.StringAttribute{Optional: true, Description: "Destination IP address or CIDR."},
						"src_port": schema.Int64Attribute{Optional: true, Description: "Source port number."},
						"dst_port": schema.Int64Attribute{Optional: true, Description: "Destination port number."},
					},
				},
			},
		},
	}
}
