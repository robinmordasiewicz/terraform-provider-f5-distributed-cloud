// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package network_firewall

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Network Firewall.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_network_firewall`" + ` resource manages Network Firewalls in F5 Distributed Cloud.

Network Firewalls provide network-level security controls.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_network_firewall" "example" {
  name        = "my-network-firewall"
  namespace   = "system"
  description = "Network firewall for perimeter security"
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
