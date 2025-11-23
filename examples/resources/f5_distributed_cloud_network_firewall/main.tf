# Example configuration for f5_distributed_cloud_network_firewall

resource "f5_distributed_cloud_network_firewall" "example" {
  name        = "example-network_firewall"
  namespace   = "system"
  description = "Example NetworkFirewall resource managed by Terraform"

  # Add additional configuration as needed
}
