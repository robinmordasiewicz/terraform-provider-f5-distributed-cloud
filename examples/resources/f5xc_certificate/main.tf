# Example configuration for f5xc_certificate

resource "f5xc_certificate" "example" {
  name        = "example-certificate"
  namespace   = "system"
  description = "Example Certificate resource managed by Terraform"

  # Add additional configuration as needed
}
