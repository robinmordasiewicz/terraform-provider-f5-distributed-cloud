# Example configuration for f5xc_user_group

resource "f5xc_user_group" "example" {
  name        = "example-user_group"
  namespace   = "system"
  description = "Example UserGroup resource managed by Terraform"

  # Add additional configuration as needed
}
