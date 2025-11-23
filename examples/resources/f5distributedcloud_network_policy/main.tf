# Example configuration for f5distributedcloud_network_policy

resource "f5distributedcloud_network_policy" "example" {
  name        = "example-network_policy"
  namespace   = "system"
  description = "Example NetworkPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
