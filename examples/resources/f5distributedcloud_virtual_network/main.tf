# Example configuration for f5distributedcloud_virtual_network

resource "f5distributedcloud_virtual_network" "example" {
  name        = "example-virtual_network"
  namespace   = "system"
  description = "Example VirtualNetwork resource managed by Terraform"

  # Add additional configuration as needed
}
