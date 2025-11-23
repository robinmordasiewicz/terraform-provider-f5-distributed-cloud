# Example configuration for f5_distributed_cloud_tcp_loadbalancer

resource "f5_distributed_cloud_tcp_loadbalancer" "example" {
  name        = "example-tcp_loadbalancer"
  namespace   = "system"
  description = "Example TCPLoadbalancer resource managed by Terraform"

  # Add additional configuration as needed
}
