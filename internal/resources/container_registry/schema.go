// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package container_registry

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Container Registry resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Container Registry. Container registries store and manage container images.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_container_registry`" + ` resource manages Container Registries in F5 Distributed Cloud.

Container registries store and manage container images.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_container_registry" "example" {
  name        = "my-container-registry"
  namespace   = "shared"
  description = "Container registry for application images"
}
` + "```" + `

## Import

Container Registries can be imported using namespace/name:

` + "```shell" + `
terraform import f5_distributed_cloud_container_registry.example shared/my-container-registry
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Container Registry.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Container Registry.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Container Registry is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Container Registry.",
				Optional:    true,
			},
		},
	}
}
