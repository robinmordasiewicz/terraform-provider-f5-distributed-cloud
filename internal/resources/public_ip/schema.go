// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package public_ip

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Public IP resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Public IP. Public IPs provide external IP addresses for services.",
		MarkdownDescription: `
The ` + "`f5xc_public_ip`" + ` resource manages Public IPs in F5 Distributed Cloud.

Public IPs provide external IP addresses that can be used for services deployed
within the F5 Distributed Cloud platform.

## Example Usage

` + "```hcl" + `
resource "f5xc_public_ip" "example" {
  name        = "my-public-ip"
  namespace   = "shared"
  description = "Public IP for production services"
}
` + "```" + `

## Import

Public IPs can be imported using namespace/name:

` + "```shell" + `
terraform import f5xc_public_ip.example shared/my-public-ip
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Public IP.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Public IP.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Public IP is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Public IP.",
				Optional:    true,
			},
		},
	}
}
