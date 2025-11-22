# Example configuration for f5xc_role

resource "f5xc_role" "example" {
  name        = "example-role"
  namespace   = "system"
  description = "Example Role resource managed by Terraform"

  # Add additional configuration as needed
}
