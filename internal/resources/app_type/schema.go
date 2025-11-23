// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_type

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud App Type.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_app_type`" + ` resource manages App Types in F5 Distributed Cloud.

App Types define application characteristics and enable features like API discovery, sensitive data detection, and more.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_app_type" "example" {
  name      = "my-app-type"
  namespace = "shared"

  features {
    type = "USER_IDENTIFICATION"
  }

  features {
    type = "API_DISCOVERY"
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
			"features": schema.ListNestedBlock{
				Description: "List of features to enable for the app type.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{Required: true, Description: "Feature type (e.g., USER_IDENTIFICATION, API_DISCOVERY)."},
					},
				},
			},
		},
	}
}
