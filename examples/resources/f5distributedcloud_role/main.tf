# Example configuration for f5distributedcloud_role

resource "f5distributedcloud_role" "example" {
  name        = "example-role"
  namespace   = "system"
  description = "Example Role resource managed by Terraform"

  # Add additional configuration as needed
}
