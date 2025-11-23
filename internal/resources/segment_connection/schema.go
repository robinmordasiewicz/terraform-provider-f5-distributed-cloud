// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package segment_connection

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Segment Connection resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Segment Connection. Segment Connections link network segments together.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_segment_connection`" + ` resource manages Segment Connections in F5 Distributed Cloud.

Segment Connections link network segments together, enabling communication between
different network segments within the F5 Distributed Cloud infrastructure.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_segment_connection" "example" {
  name        = "my-segment-connection"
  namespace   = "system"
  description = "Connection between production and staging segments"

  segment1 {
    name      = "production-segment"
    namespace = "system"
  }

  segment2 {
    name      = "staging-segment"
    namespace = "system"
  }
}
` + "```" + `

## Import

Segment Connections can be imported using namespace/name:

` + "```shell" + `
terraform import f5distributedcloud_segment_connection.example system/my-segment-connection
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Segment Connection.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Segment Connection.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Segment Connection is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Segment Connection.",
				Optional:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"segment1": schema.SingleNestedBlock{
				Description: "Reference to the first network segment.",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "Name of the segment.",
						Required:    true,
					},
					"namespace": schema.StringAttribute{
						Description: "Namespace of the segment.",
						Required:    true,
					},
				},
			},
			"segment2": schema.SingleNestedBlock{
				Description: "Reference to the second network segment.",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "Name of the segment.",
						Required:    true,
					},
					"namespace": schema.StringAttribute{
						Description: "Namespace of the segment.",
						Required:    true,
					},
				},
			},
		},
	}
}
