// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_site

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Schema returns the schema for the Cloud Site resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Site (AWS VPC, Azure VNET, or GCP VPC site).",
		MarkdownDescription: `
The ` + "`f5distributedcloud_cloud_site`" + ` resource manages Cloud Sites in F5 Distributed Cloud.

Cloud Sites are deployed infrastructure in public cloud providers that enable
F5 Distributed Cloud services such as load balancing, security, and networking.

## Supported Cloud Providers

- **AWS**: Creates VPC Sites in Amazon Web Services
- **Azure**: Creates VNET Sites in Microsoft Azure
- **GCP**: Creates VPC Sites in Google Cloud Platform

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_cloud_site" "aws_site" {
  name           = "my-aws-site"
  description    = "AWS VPC Site"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-east-1"

  vpc_name = "my-vpc"
  cidr     = "10.0.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.0.1.0/24"
    availability_zone = "us-east-1a"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "t3.xlarge"
    disk_size     = 80
    count         = 1
  }
}
` + "```" + `

## Import

Cloud Sites can be imported using the site name:

` + "```shell" + `
terraform import f5distributedcloud_cloud_site.example my-aws-site
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Cloud Site.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Cloud Site.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Cloud Site.",
				Optional:    true,
			},
			"site_type": schema.StringAttribute{
				Description: "Type of cloud site: 'aws_vpc_site', 'azure_vnet_site', or 'gcp_vpc_site'.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("aws_vpc_site", "azure_vnet_site", "gcp_vpc_site"),
				},
			},
			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider: 'aws', 'azure', or 'gcp'.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("aws", "azure", "gcp"),
				},
			},
			"region": schema.StringAttribute{
				Description: "Cloud provider region for the site.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vpc_name": schema.StringAttribute{
				Description: "Name of the VPC/VNET to create or use.",
				Optional:    true,
			},
			"vpc_id": schema.StringAttribute{
				Description: "ID of an existing VPC/VNET to use (for existing VPC deployments).",
				Optional:    true,
			},
			"cidr": schema.StringAttribute{
				Description: "CIDR block for the VPC/VNET.",
				Optional:    true,
			},
			"ssh_key": schema.StringAttribute{
				Description: "SSH public key for node access.",
				Optional:    true,
				Sensitive:   true,
			},
			"labels": schema.MapAttribute{
				Description: "Labels to apply to the site.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"annotations": schema.MapAttribute{
				Description: "Annotations for the site.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"site_state": schema.StringAttribute{
				Description: "Current state of the site (computed).",
				Computed:    true,
			},
			"software_version": schema.StringAttribute{
				Description: "Software version for the site nodes.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
			},
			"operating_system": schema.StringAttribute{
				Description: "Operating system for site nodes.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
			},
		},
		Blocks: map[string]schema.Block{
			"subnets": schema.ListNestedBlock{
				Description: "Subnet configurations for the site.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of the subnet.",
							Required:    true,
						},
						"cidr": schema.StringAttribute{
							Description: "CIDR block for the subnet.",
							Required:    true,
						},
						"availability_zone": schema.StringAttribute{
							Description: "Availability zone for the subnet.",
							Required:    true,
						},
						"subnet_type": schema.StringAttribute{
							Description: "Type of subnet: 'public' or 'private'.",
							Optional:    true,
							Computed:    true,
							Default:     stringdefault.StaticString("private"),
							Validators: []validator.String{
								stringvalidator.OneOf("public", "private"),
							},
						},
					},
				},
			},
			"master_nodes": schema.SingleNestedBlock{
				Description: "Master node configuration.",
				Attributes: map[string]schema.Attribute{
					"instance_type": schema.StringAttribute{
						Description: "Instance type for master nodes.",
						Required:    true,
					},
					"disk_size": schema.Int64Attribute{
						Description: "Disk size in GB for master nodes.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(80),
					},
					"count": schema.Int64Attribute{
						Description: "Number of master nodes (1 or 3 for HA).",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1),
					},
				},
			},
			"worker_nodes": schema.SingleNestedBlock{
				Description: "Worker node configuration for sites with worker pools.",
				Attributes: map[string]schema.Attribute{
					"instance_type": schema.StringAttribute{
						Description: "Instance type for worker nodes.",
						Required:    true,
					},
					"disk_size": schema.Int64Attribute{
						Description: "Disk size in GB for worker nodes.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(80),
					},
					"count": schema.Int64Attribute{
						Description: "Number of worker nodes.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
					},
					"min_count": schema.Int64Attribute{
						Description: "Minimum number of worker nodes for autoscaling.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
					},
					"max_count": schema.Int64Attribute{
						Description: "Maximum number of worker nodes for autoscaling.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
					},
				},
			},
		},
	}
}
