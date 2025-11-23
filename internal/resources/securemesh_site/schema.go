// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package securemesh_site

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Securemesh Site.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_securemesh_site`" + ` resource manages Securemesh Sites in F5 Distributed Cloud.

Securemesh Sites provide secure connectivity for distributed applications.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_securemesh_site" "example" {
  name        = "my-securemesh-site"
  namespace   = "system"
  description = "Example securemesh site"
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
