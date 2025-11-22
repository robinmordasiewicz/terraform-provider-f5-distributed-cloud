# Example configuration for f5xc_registration

resource "f5xc_registration" "example" {
  name        = "example-registration"
  namespace   = "system"
  description = "Example Registration resource managed by Terraform"

  # Add additional configuration as needed
}
