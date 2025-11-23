# Example configuration for f5_distributed_cloud_network_policy

resource "f5_distributed_cloud_network_policy" "example" {
  name        = "example-network_policy"
  namespace   = "system"
  description = "Example NetworkPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
