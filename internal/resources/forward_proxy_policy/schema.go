// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package forward_proxy_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Forward Proxy Policy.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_forward_proxy_policy`" + ` resource manages Forward Proxy Policies in F5 Distributed Cloud.

Forward Proxy Policies define rules for controlling outbound traffic from applications through the F5 XC proxy.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_forward_proxy_policy" "example" {
  name      = "my-proxy-policy"
  namespace = "my-namespace"

  proxy_rules {
    action       = "allow"
    domain_match = "*.example.com"
    port_match   = 443
    http_methods = ["GET", "POST"]
  }

  proxy_rules {
    action       = "deny"
    domain_match = "*.blocked.com"
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
			"proxy_rules": schema.ListNestedBlock{
				Description: "Proxy rule configuration.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Required:    true,
							Description: "Action to take (allow/deny).",
							Validators:  []validator.String{stringvalidator.OneOf("allow", "deny")},
						},
						"domain_match": schema.StringAttribute{
							Optional:    true,
							Description: "Domain pattern to match (supports wildcards).",
						},
						"port_match": schema.Int64Attribute{
							Optional:    true,
							Description: "Port number to match.",
						},
						"http_methods": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
							Description: "HTTP methods to match.",
						},
					},
				},
			},
		},
	}
}
