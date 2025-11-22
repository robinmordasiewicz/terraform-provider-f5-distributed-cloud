// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package voltstack_site

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Voltstack Site.",
		MarkdownDescription: `
The ` + "`f5xc_voltstack_site`" + ` resource manages Voltstack Sites in F5 Distributed Cloud.

Voltstack Sites are F5-managed compute infrastructure deployed at the edge.

## Example Usage

` + "```hcl" + `
resource "f5xc_voltstack_site" "example" {
  name           = "edge-site-1"
  namespace      = "system"
  master_nodes   = 3
  worker_nodes   = 2
  site_type      = "REGIONAL_EDGE"
  volterra_region = "ves-io-us-west-2"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"master_nodes": schema.Int64Attribute{
				Required:    true,
				Description: "Number of master nodes.",
			},
			"worker_nodes": schema.Int64Attribute{
				Optional:    true,
				Description: "Number of worker nodes.",
			},
			"site_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of Voltstack site.",
			},
			"volterra_region": schema.StringAttribute{
				Optional:    true,
				Description: "Volterra region for the site.",
			},
		},
	}
}
