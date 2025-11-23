# Example configuration for f5distributedcloud_policer

resource "f5distributedcloud_policer" "example" {
  name        = "example-policer"
  namespace   = "system"
  description = "Example Policer resource managed by Terraform"

  # Add additional configuration as needed
}
