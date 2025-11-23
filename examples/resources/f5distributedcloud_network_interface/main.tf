# Example configuration for f5distributedcloud_network_interface

resource "f5distributedcloud_network_interface" "example" {
  name        = "example-network_interface"
  namespace   = "system"
  description = "Example NetworkInterface resource managed by Terraform"

  # Add additional configuration as needed
}
