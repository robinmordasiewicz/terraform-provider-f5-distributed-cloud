# Example configuration for f5distributedcloud_aws_vpc_site

resource "f5distributedcloud_aws_vpc_site" "example" {
  name        = "example-aws_vpc_site"
  namespace   = "system"
  description = "Example AWSVPCSite resource managed by Terraform"

  # Add additional configuration as needed
}
