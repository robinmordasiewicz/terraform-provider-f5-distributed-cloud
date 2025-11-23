// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_mesh

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Service Mesh.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_service_mesh`" + ` resource manages Service Meshes in F5 Distributed Cloud.

Service Meshes provide service-to-service communication management.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_service_mesh" "example" {
  name      = "app-mesh"
  namespace = "my-namespace"
  mesh_type = "INGRESS"
  enabled   = true
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"mesh_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of service mesh (INGRESS, EGRESS, SIDECAR).",
				Validators:  []validator.String{stringvalidator.OneOf("INGRESS", "EGRESS", "SIDECAR")},
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the service mesh is enabled.",
			},
		},
	}
}
