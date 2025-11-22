---
page_title: "f5xc_cloud_site Resource - F5 Distributed Cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud Site (AWS VPC, Azure VNET, or GCP VPC).
---

# f5xc_cloud_site (Resource)

Manages an F5 Distributed Cloud Site in AWS, Azure, or GCP.

Cloud Sites are deployed infrastructure in public cloud providers that enable F5 Distributed Cloud services.

## Example Usage

### AWS VPC Site

```terraform
resource "f5xc_cloud_site" "aws" {
  name           = "aws-site"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-east-1"

  vpc_name = "production-vpc"
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
```

### Azure VNET Site

```terraform
resource "f5xc_cloud_site" "azure" {
  name           = "azure-site"
  site_type      = "azure_vnet_site"
  cloud_provider = "azure"
  region         = "eastus"

  vpc_name = "production-vnet"
  cidr     = "10.0.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.0.1.0/24"
    availability_zone = "1"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "Standard_D3_v2"
    disk_size     = 80
    count         = 1
  }
}
```

## Argument Reference

- `name` - (Required) Name of the cloud site. Changing this forces a new resource.
- `site_type` - (Required) Type of site: `aws_vpc_site`, `azure_vnet_site`, or `gcp_vpc_site`.
- `cloud_provider` - (Required) Cloud provider: `aws`, `azure`, or `gcp`.
- `region` - (Required) Cloud provider region.
- `description` - (Optional) Description of the site.
- `vpc_name` - (Optional) Name of the VPC/VNET.
- `vpc_id` - (Optional) ID of existing VPC/VNET.
- `cidr` - (Optional) CIDR block for the VPC/VNET.
- `ssh_key` - (Optional) SSH public key for node access.
- `labels` - (Optional) Labels for the site.
- `subnets` - (Optional) Subnet configurations.
- `master_nodes` - (Optional) Master node configuration.
- `worker_nodes` - (Optional) Worker node configuration.

### subnets

- `name` - (Required) Name of the subnet.
- `cidr` - (Required) CIDR block.
- `availability_zone` - (Required) Availability zone.
- `subnet_type` - (Optional) Type: `public` or `private`. Defaults to `private`.

### master_nodes / worker_nodes

- `instance_type` - (Required) Instance type.
- `disk_size` - (Optional) Disk size in GB. Defaults to 80.
- `count` - (Optional) Number of nodes.

## Attribute Reference

- `id` - The unique identifier of the cloud site.
- `site_state` - Current state of the site.

## Import

Cloud Sites can be imported using the name:

```shell
terraform import f5xc_cloud_site.example aws-site
```
