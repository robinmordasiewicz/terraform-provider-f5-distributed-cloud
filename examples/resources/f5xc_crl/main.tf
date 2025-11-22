# Example configuration for f5xc_crl

resource "f5xc_crl" "example" {
  name        = "example-crl"
  namespace   = "system"
  description = "Example CRL resource managed by Terraform"

  # Add additional configuration as needed
}
