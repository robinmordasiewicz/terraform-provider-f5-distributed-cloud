// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_host

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Virtual Host.",
		MarkdownDescription: `
The ` + "`f5xc_virtual_host`" + ` resource manages Virtual Hosts in F5 Distributed Cloud.

Virtual Hosts define HTTP routing configurations for domain-based traffic.

## Example Usage

` + "```hcl" + `
resource "f5xc_virtual_host" "example" {
  name        = "my-vhost"
  namespace   = "system"
  description = "Virtual host for web application"
  domains     = ["app.example.com", "www.example.com"]
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"domains":     schema.ListAttribute{Optional: true, ElementType: types.StringType, Description: "List of domains for this virtual host."},
		},
	}
}
