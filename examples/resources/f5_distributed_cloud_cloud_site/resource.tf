# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: AWS VPC Site
resource "f5_distributed_cloud_cloud_site" "aws_site" {
  name           = "my-aws-site"
  description    = "AWS VPC Site for production workloads"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-east-1"

  vpc_name = "production-vpc"
  cidr     = "10.0.0.0/16"

  subnets {
    name              = "public-subnet-1"
    cidr              = "10.0.1.0/24"
    availability_zone = "us-east-1a"
    subnet_type       = "public"
  }

  subnets {
    name              = "private-subnet-1"
    cidr              = "10.0.10.0/24"
    availability_zone = "us-east-1a"
    subnet_type       = "private"
  }

  master_nodes {
    instance_type = "t3.xlarge"
    disk_size     = 100
    count         = 1
  }

  labels = {
    environment = "production"
    team        = "platform"
  }
}

# Example: AWS VPC Site with High Availability
resource "f5_distributed_cloud_cloud_site" "aws_ha_site" {
  name           = "my-aws-ha-site"
  description    = "AWS VPC Site with HA configuration"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-west-2"

  vpc_name = "ha-vpc"
  cidr     = "10.1.0.0/16"

  subnets {
    name              = "subnet-az-a"
    cidr              = "10.1.1.0/24"
    availability_zone = "us-west-2a"
    subnet_type       = "public"
  }

  subnets {
    name              = "subnet-az-b"
    cidr              = "10.1.2.0/24"
    availability_zone = "us-west-2b"
    subnet_type       = "public"
  }

  subnets {
    name              = "subnet-az-c"
    cidr              = "10.1.3.0/24"
    availability_zone = "us-west-2c"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "t3.xlarge"
    disk_size     = 100
    count         = 3  # HA configuration
  }
}

# Example: Azure VNET Site
resource "f5_distributed_cloud_cloud_site" "azure_site" {
  name           = "my-azure-site"
  description    = "Azure VNET Site"
  site_type      = "azure_vnet_site"
  cloud_provider = "azure"
  region         = "eastus"

  vpc_name = "production-vnet"
  cidr     = "10.2.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.2.1.0/24"
    availability_zone = "1"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "Standard_D3_v2"
    disk_size     = 80
    count         = 1
  }
}

# Example: GCP VPC Site
resource "f5_distributed_cloud_cloud_site" "gcp_site" {
  name           = "my-gcp-site"
  description    = "GCP VPC Site"
  site_type      = "gcp_vpc_site"
  cloud_provider = "gcp"
  region         = "us-central1"

  vpc_name = "production-vpc"
  cidr     = "10.3.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.3.1.0/24"
    availability_zone = "us-central1-a"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "n1-standard-4"
    disk_size     = 80
    count         = 1
  }
}

# Example: AWS Site with Worker Nodes
resource "f5_distributed_cloud_cloud_site" "aws_with_workers" {
  name           = "my-aws-site-workers"
  description    = "AWS VPC Site with worker node pool"
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "eu-west-1"

  vpc_name = "worker-vpc"
  cidr     = "10.4.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.4.1.0/24"
    availability_zone = "eu-west-1a"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "t3.xlarge"
    disk_size     = 80
    count         = 1
  }

  worker_nodes {
    instance_type = "t3.large"
    disk_size     = 100
    count         = 2
    min_count     = 1
    max_count     = 5
  }
}
