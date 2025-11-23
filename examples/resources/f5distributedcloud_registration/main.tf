# Example configuration for f5distributedcloud_registration

resource "f5distributedcloud_registration" "example" {
  name        = "example-registration"
  namespace   = "system"
  description = "Example Registration resource managed by Terraform"

  # Add additional configuration as needed
}
