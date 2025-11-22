// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package gcp_vpc_site

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud GCP VPC Site.",
		MarkdownDescription: `
The ` + "`f5xc_gcp_vpc_site`" + ` resource manages GCP VPC Sites in F5 Distributed Cloud.

GCP VPC Sites deploy F5 XC nodes into an existing or new GCP VPC network.

## Example Usage

` + "```hcl" + `
resource "f5xc_gcp_vpc_site" "example" {
  name                  = "gcp-site-1"
  namespace             = "system"
  gcp_region            = "us-west1"
  project_id            = "my-gcp-project"
  network_name          = "my-vpc-network"
  machine_type          = "n1-standard-4"
  cloud_credentials_ref = "gcp-creds"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"gcp_region": schema.StringAttribute{
				Required:    true,
				Description: "GCP region for the site.",
			},
			"project_id": schema.StringAttribute{
				Required:    true,
				Description: "GCP project ID.",
			},
			"network_name": schema.StringAttribute{
				Optional:    true,
				Description: "GCP VPC network name.",
			},
			"machine_type": schema.StringAttribute{
				Optional:    true,
				Description: "GCP VM machine type for site nodes.",
			},
			"cloud_credentials_ref": schema.StringAttribute{
				Required:    true,
				Description: "Reference to cloud credentials.",
			},
		},
	}
}
