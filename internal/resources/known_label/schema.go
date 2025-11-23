// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package known_label

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Known Label.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_known_label`" + ` resource manages Known Labels in F5 Distributed Cloud.

Known Labels define predefined labels for resource categorization and selection.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_known_label" "example" {
  name      = "environment-label"
  namespace = "shared"
  key       = "environment"
  value     = "production"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"key": schema.StringAttribute{
				Required:    true,
				Description: "Label key.",
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "Label value.",
			},
		},
	}
}
