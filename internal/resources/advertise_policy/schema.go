// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package advertise_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Advertise Policy.",
		MarkdownDescription: `
The ` + "`f5xc_advertise_policy`" + ` resource manages Advertise Policies in F5 Distributed Cloud.

Advertise Policies control how services are advertised to the network.

## Example Usage

` + "```hcl" + `
resource "f5xc_advertise_policy" "example" {
  name        = "my-advertise-policy"
  namespace   = "my-namespace"
  description = "Policy for service advertisement"
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
