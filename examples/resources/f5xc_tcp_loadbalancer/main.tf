# Example configuration for f5xc_tcp_loadbalancer

resource "f5xc_tcp_loadbalancer" "example" {
  name        = "example-tcp_loadbalancer"
  namespace   = "system"
  description = "Example TCPLoadbalancer resource managed by Terraform"

  # Add additional configuration as needed
}
