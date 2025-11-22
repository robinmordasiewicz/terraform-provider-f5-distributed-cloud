# Example configuration for f5xc_contact

resource "f5xc_contact" "example" {
  name        = "example-contact"
  namespace   = "system"
  description = "Example Contact resource managed by Terraform"

  # Add additional configuration as needed
}
