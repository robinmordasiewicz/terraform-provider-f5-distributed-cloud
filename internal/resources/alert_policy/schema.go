// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package alert_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Alert Policy.",
		MarkdownDescription: `
The ` + "`f5xc_alert_policy`" + ` resource manages Alert Policies in F5 Distributed Cloud.

Alert Policies define notification rules for alerts and events.

## Example Usage

` + "```hcl" + `
resource "f5xc_alert_policy" "example" {
  name        = "my-alert-policy"
  namespace   = "my-namespace"
  description = "Alert policy for critical alerts"

  receivers {
    name      = "slack-receiver"
    namespace = "shared"
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
			"receivers": schema.ListNestedBlock{
				Description: "List of alert receivers.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name":      schema.StringAttribute{Required: true},
						"namespace": schema.StringAttribute{Required: true},
					},
				},
			},
			"routes": schema.ListNestedBlock{
				Description: "Alert routing rules.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"receiver": schema.StringAttribute{Required: true},
						"match":    schema.MapAttribute{Optional: true, ElementType: types.StringType},
					},
				},
			},
		},
	}
}
