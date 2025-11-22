# Example configuration for f5xc_aws_tgw_site

resource "f5xc_aws_tgw_site" "example" {
  name        = "example-aws_tgw_site"
  namespace   = "system"
  description = "Example AWSTGWSite resource managed by Terraform"

  # Add additional configuration as needed
}
