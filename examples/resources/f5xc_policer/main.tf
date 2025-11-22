# Example configuration for f5xc_policer

resource "f5xc_policer" "example" {
  name        = "example-policer"
  namespace   = "system"
  description = "Example Policer resource managed by Terraform"

  # Add additional configuration as needed
}
