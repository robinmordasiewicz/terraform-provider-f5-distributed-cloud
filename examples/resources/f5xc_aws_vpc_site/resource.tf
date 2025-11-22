# Example: AWS VPC Site
# This example creates an AWS VPC Site in F5 Distributed Cloud

# First, ensure you have cloud credentials configured
resource "f5xc_cloud_credentials" "aws" {
  name      = "aws-creds"
  namespace = "system"

  aws_secret_key {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
  }
}

# Create the AWS VPC Site
resource "f5xc_aws_vpc_site" "example" {
  name                  = "my-aws-site"
  namespace             = "system"
  description           = "Example AWS VPC Site"
  aws_region            = "us-west-2"
  vpc_id                = "vpc-0123456789abcdef0"
  instance_type         = "t3.xlarge"
  cloud_credentials_ref = f5xc_cloud_credentials.aws.name
}

# Variables
variable "aws_access_key" {
  description = "AWS Access Key"
  type        = string
  sensitive   = true
}

variable "aws_secret_key" {
  description = "AWS Secret Key"
  type        = string
  sensitive   = true
}

# Outputs
output "site_name" {
  description = "Name of the created AWS VPC Site"
  value       = f5xc_aws_vpc_site.example.name
}

output "site_id" {
  description = "ID of the created AWS VPC Site"
  value       = f5xc_aws_vpc_site.example.id
}
