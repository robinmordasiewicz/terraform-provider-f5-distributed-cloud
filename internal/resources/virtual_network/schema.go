// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_network

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Virtual Network.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_virtual_network`" + ` resource manages Virtual Networks in F5 Distributed Cloud.

Virtual Networks provide isolated network environments for workloads.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_virtual_network" "example" {
  name        = "my-vnet"
  namespace   = "system"
  description = "Virtual network for application isolation"
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
