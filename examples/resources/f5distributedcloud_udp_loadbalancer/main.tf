# Example configuration for f5distributedcloud_udp_loadbalancer

resource "f5distributedcloud_udp_loadbalancer" "example" {
  name        = "example-udp_loadbalancer"
  namespace   = "system"
  description = "Example UDPLoadbalancer resource managed by Terraform"

  # Add additional configuration as needed
}
