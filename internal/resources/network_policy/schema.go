// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package network_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Network Policy.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_network_policy`" + ` resource manages Network Policies in F5 Distributed Cloud.

Network Policies define traffic rules for controlling network connectivity between workloads.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_network_policy" "example" {
  name        = "network-policy"
  namespace   = "my-namespace"
  policy_type = "INGRESS"

  rules {
    name          = "allow-web"
    action        = "ALLOW"
    source_prefix = "10.0.0.0/8"
    protocol      = "TCP"
    ports         = "80,443"
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"policy_type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of network policy (INGRESS/EGRESS).",
				Validators:  []validator.String{stringvalidator.OneOf("INGRESS", "EGRESS")},
			},
		},
		Blocks: map[string]schema.Block{
			"rules": schema.ListNestedBlock{
				Description: "Network policy rules.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{Required: true, Description: "Name of the rule."},
						"action": schema.StringAttribute{
							Required:    true,
							Description: "Action to take (ALLOW/DENY).",
							Validators:  []validator.String{stringvalidator.OneOf("ALLOW", "DENY")},
						},
						"source_prefix": schema.StringAttribute{Optional: true, Description: "Source IP prefix (CIDR)."},
						"dest_prefix":   schema.StringAttribute{Optional: true, Description: "Destination IP prefix (CIDR)."},
						"protocol": schema.StringAttribute{
							Optional:    true,
							Description: "Protocol (TCP/UDP/ICMP).",
							Validators:  []validator.String{stringvalidator.OneOf("TCP", "UDP", "ICMP", "ANY")},
						},
						"ports": schema.StringAttribute{Optional: true, Description: "Port numbers (comma-separated or range)."},
					},
				},
			},
		},
	}
}
