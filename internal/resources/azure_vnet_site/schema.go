// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package azure_vnet_site

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Azure VNET Site.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_azure_vnet_site`" + ` resource manages Azure VNET Sites in F5 Distributed Cloud.

Azure VNET Sites deploy F5 XC nodes into an existing or new Azure VNET.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_azure_vnet_site" "example" {
  name                  = "azure-site-1"
  namespace             = "system"
  azure_region          = "westus2"
  resource_group        = "my-resource-group"
  vnet_name             = "my-vnet"
  machine_type          = "Standard_D3_v2"
  cloud_credentials_ref = "azure-creds"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"azure_region": schema.StringAttribute{
				Required:    true,
				Description: "Azure region for the site.",
			},
			"resource_group": schema.StringAttribute{
				Required:    true,
				Description: "Azure resource group name.",
			},
			"vnet_name": schema.StringAttribute{
				Optional:    true,
				Description: "Azure VNET name.",
			},
			"machine_type": schema.StringAttribute{
				Optional:    true,
				Description: "Azure VM size for site nodes.",
			},
			"cloud_credentials_ref": schema.StringAttribute{
				Required:    true,
				Description: "Reference to cloud credentials.",
			},
		},
	}
}
