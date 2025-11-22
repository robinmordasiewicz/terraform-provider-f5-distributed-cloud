// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package workload

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Workload resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Workload. Workloads represent application deployments.",
		MarkdownDescription: `
The ` + "`f5xc_workload`" + ` resource manages Workloads in F5 Distributed Cloud.

Workloads represent application deployments that can be managed and monitored
within the F5 Distributed Cloud platform.

## Example Usage

` + "```hcl" + `
resource "f5xc_workload" "example" {
  name        = "my-workload"
  namespace   = "my-namespace"
  description = "Example workload for my application"
}
` + "```" + `

## Import

Workloads can be imported using namespace/name:

` + "```shell" + `
terraform import f5xc_workload.example my-namespace/my-workload
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Workload.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Workload. Must be unique within the namespace and follow F5 XC naming conventions.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Workload is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Workload.",
				Optional:    true,
			},
		},
	}
}
