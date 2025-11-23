# Example configuration for f5distributedcloud_tunnel

resource "f5distributedcloud_tunnel" "example" {
  name        = "example-tunnel"
  namespace   = "system"
  description = "Example Tunnel resource managed by Terraform"

  # Add additional configuration as needed
}
