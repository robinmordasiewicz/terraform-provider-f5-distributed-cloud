// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_policy_set

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Service Policy Set. Service Policy Sets group service policies together.",
		MarkdownDescription: `
The ` + "`f5xc_service_policy_set`" + ` resource manages Service Policy Sets in F5 Distributed Cloud.

Service Policy Sets group service policies together for organization and management purposes.

## Example Usage

` + "```hcl" + `
resource "f5xc_service_policy_set" "example" {
  name        = "my-policy-set"
  namespace   = "my-namespace"
  description = "Example service policy set"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the service policy set.",
				MarkdownDescription: "The unique identifier of the service policy set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the service policy set.",
				MarkdownDescription: "The name of the service policy set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Required:            true,
				Description:         "The namespace where the service policy set will be created.",
				MarkdownDescription: "The namespace where the service policy set will be created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Description:         "A description of the service policy set.",
				MarkdownDescription: "A description of the service policy set.",
			},
		},
	}
}
