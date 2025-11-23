// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_site

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Schema returns the schema for the Virtual Site resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Virtual Site.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_virtual_site`" + ` resource manages Virtual Sites in F5 Distributed Cloud.

Virtual Sites are logical groupings of physical sites that allow you to apply
configurations and policies to multiple sites as a single entity.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_virtual_site" "example" {
  name        = "my-virtual-site"
  namespace   = "shared"
  description = "Virtual site for production"

  site_type = "REGIONAL_EDGE"

  site_selector {
    expressions = ["site.region == 'us-east'"]
  }
}
` + "```" + `

## Import

Virtual Sites can be imported using namespace/name:

` + "```shell" + `
terraform import f5distributedcloud_virtual_site.example shared/my-virtual-site
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Virtual Site.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Virtual Site.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Virtual Site is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Virtual Site.",
				Optional:    true,
			},
			"site_type": schema.StringAttribute{
				Description: "Type of sites to include: 'REGIONAL_EDGE', 'CUSTOMER_EDGE'.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("REGIONAL_EDGE", "CUSTOMER_EDGE"),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"site_selector": schema.SingleNestedBlock{
				Description: "Selector for matching sites to include in the virtual site.",
				Attributes: map[string]schema.Attribute{
					"expressions": schema.ListAttribute{
						Description: "List of label selector expressions.",
						Required:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}
