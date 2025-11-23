// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cdn_cache_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud CDN Cache Rule.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_cdn_cache_rule`" + ` resource manages CDN Cache Rules in F5 Distributed Cloud.

CDN Cache Rules define caching behavior for specific paths or content types.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_cdn_cache_rule" "example" {
  name        = "static-assets-cache"
  namespace   = "system"
  description = "Cache static assets for 1 hour"
  path_match  = "/static/*"
  cache_ttl   = 3600
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"path_match":  schema.StringAttribute{Required: true, Description: "Path pattern to match for caching."},
			"cache_ttl":   schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(3600), Description: "Cache TTL in seconds."},
		},
	}
}
