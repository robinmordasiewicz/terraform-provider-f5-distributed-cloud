# Example configuration for f5distributedcloud_crl

resource "f5distributedcloud_crl" "example" {
  name        = "example-crl"
  namespace   = "system"
  description = "Example CRL resource managed by Terraform"

  # Add additional configuration as needed
}
