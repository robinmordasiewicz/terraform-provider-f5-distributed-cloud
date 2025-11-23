# Example configuration for f5distributedcloud_fleet

resource "f5distributedcloud_fleet" "example" {
  name        = "example-fleet"
  namespace   = "system"
  description = "Example Fleet resource managed by Terraform"

  # Add additional configuration as needed
}
