# Example configuration for f5xc_tunnel

resource "f5xc_tunnel" "example" {
  name        = "example-tunnel"
  namespace   = "system"
  description = "Example Tunnel resource managed by Terraform"

  # Add additional configuration as needed
}
