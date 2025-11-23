// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package site_mesh_group

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Site Mesh Group.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_site_mesh_group`" + ` resource manages Site Mesh Groups in F5 Distributed Cloud.

Site Mesh Groups enable network connectivity between sites in a service mesh topology.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_site_mesh_group" "example" {
  name      = "my-mesh-group"
  namespace = "system"
  type      = "SITE_MESH_GROUP_TYPE_FULL_MESH"

  virtual_site {
    name      = "my-virtual-site"
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
			"type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of site mesh group (SITE_MESH_GROUP_TYPE_FULL_MESH, SITE_MESH_GROUP_TYPE_HUB_SPOKE).",
				Validators:  []validator.String{stringvalidator.OneOf("SITE_MESH_GROUP_TYPE_FULL_MESH", "SITE_MESH_GROUP_TYPE_HUB_SPOKE")},
			},
		},
		Blocks: map[string]schema.Block{
			"virtual_site": schema.ListNestedBlock{
				Description: "Virtual sites to include in the mesh group.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name":      schema.StringAttribute{Required: true, Description: "Name of the virtual site."},
						"namespace": schema.StringAttribute{Required: true, Description: "Namespace of the virtual site."},
					},
				},
			},
		},
	}
}
