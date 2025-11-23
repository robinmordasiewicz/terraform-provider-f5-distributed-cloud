// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/validators"
)

// SchemaVersion is the current schema version for the namespace resource.
const SchemaVersion = 1

// Schema returns the schema for the namespace resource.
func Schema() schema.Schema {
	return schema.Schema{
		Version:     SchemaVersion,
		Description: "Manages an F5 Distributed Cloud namespace.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_namespace`" + ` resource manages namespaces in F5 Distributed Cloud.

Namespaces are the fundamental organizational unit in F5 XC, providing logical isolation
for resources like load balancers, origin pools, and application firewalls.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_namespace" "example" {
  name        = "my-namespace"
  description = "Example namespace for my application"

  timeouts {
    create = "10m"
    update = "10m"
    delete = "10m"
  }
}
` + "```" + `

## Import

Namespaces can be imported using their name:

` + "```shell" + `
terraform import f5_distributed_cloud_namespace.example my-namespace
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the namespace.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the namespace. Must be unique and follow F5 XC naming conventions (lowercase letters, numbers, hyphens; max 32 characters).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					validators.F5XCNamespaceValidator(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the namespace.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(256),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(
				nil,
				timeouts.Opts{
					Create: true,
					Read:   true,
					Update: true,
					Delete: true,
				},
			),
		},
	}
}
