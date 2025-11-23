// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud DNS Domain.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_dns_domain`" + ` resource manages DNS Domains in F5 Distributed Cloud.

DNS Domains define domain configurations for F5 XC DNS services.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_dns_domain" "example" {
  name        = "my-dns-domain"
  namespace   = "system"
  description = "Primary DNS domain"
  domain      = "app.example.com"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"domain":      schema.StringAttribute{Required: true, Description: "The domain name."},
		},
	}
}
