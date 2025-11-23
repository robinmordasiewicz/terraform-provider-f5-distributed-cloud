# Example configuration for f5_distributed_cloud_tunnel

resource "f5_distributed_cloud_tunnel" "example" {
  name        = "example-tunnel"
  namespace   = "system"
  description = "Example Tunnel resource managed by Terraform"

  # Add additional configuration as needed
}
