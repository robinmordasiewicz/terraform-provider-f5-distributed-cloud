// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Cluster resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Cluster. Clusters provide container orchestration.",
		MarkdownDescription: `
The ` + "`f5xc_cluster`" + ` resource manages Clusters in F5 Distributed Cloud.

Clusters provide container orchestration capabilities, allowing you to deploy
and manage containerized applications across your F5 Distributed Cloud infrastructure.

## Example Usage

` + "```hcl" + `
resource "f5xc_cluster" "example" {
  name        = "my-cluster"
  namespace   = "shared"
  description = "Production cluster for container workloads"
}
` + "```" + `

## Import

Clusters can be imported using namespace/name:

` + "```shell" + `
terraform import f5xc_cluster.example shared/my-cluster
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Cluster.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Cluster.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Cluster is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Cluster.",
				Optional:    true,
			},
		},
	}
}
