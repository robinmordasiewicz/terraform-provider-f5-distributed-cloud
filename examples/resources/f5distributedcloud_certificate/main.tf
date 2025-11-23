# Example configuration for f5distributedcloud_certificate

resource "f5distributedcloud_certificate" "example" {
  name        = "example-certificate"
  namespace   = "system"
  description = "Example Certificate resource managed by Terraform"

  # Add additional configuration as needed
}
