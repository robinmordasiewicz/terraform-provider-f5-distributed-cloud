# Example: Cloud Credentials Data Source
# This example retrieves information about existing Cloud Credentials

data "f5xc_cloud_credentials" "aws" {
  name      = "my-aws-credentials"
  namespace = "system"
}

# Output the cloud credentials details
output "creds_id" {
  description = "ID of the Cloud Credentials"
  value       = data.f5xc_cloud_credentials.aws.id
}

output "creds_description" {
  description = "Description of the Cloud Credentials"
  value       = data.f5xc_cloud_credentials.aws.description
}

output "creds_cloud_type" {
  description = "Type of cloud provider (aws, azure, gcp)"
  value       = data.f5xc_cloud_credentials.aws.cloud_type
}

# Use the data source to reference existing credentials for a cloud site
resource "f5xc_aws_vpc_site" "example" {
  name                  = "my-aws-site"
  namespace             = "system"
  description           = "AWS VPC Site using existing credentials"
  aws_region            = "us-west-2"
  vpc_id                = "vpc-0123456789abcdef0"
  instance_type         = "t3.xlarge"
  cloud_credentials_ref = data.f5xc_cloud_credentials.aws.name
}
