// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cdn_loadbalancer

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud CDN Load Balancer.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_cdn_loadbalancer`" + ` resource manages CDN Load Balancers in F5 Distributed Cloud.

CDN Load Balancers provide content delivery network capabilities with edge caching.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_cdn_loadbalancer" "example" {
  name        = "my-cdn-lb"
  namespace   = "system"
  description = "CDN load balancer for static content"
  domains     = ["cdn.example.com", "static.example.com"]
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"domains":     schema.ListAttribute{Optional: true, ElementType: types.StringType, Description: "List of domains for CDN delivery."},
		},
	}
}
