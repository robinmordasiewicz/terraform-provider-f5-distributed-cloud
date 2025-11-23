// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package external_connector

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud External Connector.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_external_connector`" + ` resource manages External Connectors in F5 Distributed Cloud.

External Connectors provide integration with external systems and services.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_external_connector" "example" {
  name        = "my-external-connector"
  namespace   = "my-namespace"
  description = "Example external connector"
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
