// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package network_policy_set

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Network Policy Set. Network Policy Sets group network policies together.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_network_policy_set`" + ` resource manages Network Policy Sets in F5 Distributed Cloud.

Network Policy Sets group network policies together for easier management and application to workloads.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_network_policy_set" "example" {
  name        = "my-policy-set"
  namespace   = "my-namespace"
  description = "Example network policy set"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the network policy set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the network policy set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Required:    true,
				Description: "Namespace where the network policy set will be created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Description of the network policy set.",
			},
		},
	}
}
