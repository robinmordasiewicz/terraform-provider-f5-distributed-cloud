# Example configuration for f5xc_network_firewall

resource "f5xc_network_firewall" "example" {
  name        = "example-network_firewall"
  namespace   = "system"
  description = "Example NetworkFirewall resource managed by Terraform"

  # Add additional configuration as needed
}
