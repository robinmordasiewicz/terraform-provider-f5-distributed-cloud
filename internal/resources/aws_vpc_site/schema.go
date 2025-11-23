// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package aws_vpc_site

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud AWS VPC Site.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_aws_vpc_site`" + ` resource manages AWS VPC Sites in F5 Distributed Cloud.

AWS VPC Sites deploy F5 XC nodes into an existing or new AWS VPC.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_aws_vpc_site" "example" {
  name                  = "aws-site-1"
  namespace             = "system"
  aws_region            = "us-west-2"
  vpc_id                = "vpc-12345678"
  instance_type         = "t3.xlarge"
  cloud_credentials_ref = "aws-creds"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"aws_region": schema.StringAttribute{
				Required:    true,
				Description: "AWS region for the site.",
			},
			"vpc_id": schema.StringAttribute{
				Optional:    true,
				Description: "AWS VPC ID.",
			},
			"instance_type": schema.StringAttribute{
				Optional:    true,
				Description: "EC2 instance type for site nodes.",
			},
			"cloud_credentials_ref": schema.StringAttribute{
				Required:    true,
				Description: "Reference to cloud credentials.",
			},
		},
	}
}
