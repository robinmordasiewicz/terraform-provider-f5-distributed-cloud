// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_load_balancer

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud DNS Load Balancer.",
		MarkdownDescription: `
The ` + "`f5xc_dns_load_balancer`" + ` resource manages DNS Load Balancers in F5 Distributed Cloud.

DNS Load Balancers provide global server load balancing (GSLB) capabilities.

## Example Usage

` + "```hcl" + `
resource "f5xc_dns_load_balancer" "example" {
  name              = "my-dns-lb"
  namespace         = "system"
  description       = "DNS load balancer for global distribution"
  domain            = "app.example.com"
  load_balance_type = "round_robin"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":              schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":         schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description":       schema.StringAttribute{Optional: true},
			"domain":            schema.StringAttribute{Required: true, Description: "The domain name for DNS load balancing."},
			"load_balance_type": schema.StringAttribute{Optional: true, Description: "Load balancing algorithm (round_robin, weighted, geo)."},
		},
	}
}
