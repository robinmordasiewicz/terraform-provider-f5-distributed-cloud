# Example configuration for f5distributedcloud_proxy

resource "f5distributedcloud_proxy" "example" {
  name        = "example-proxy"
  namespace   = "system"
  description = "Example Proxy resource managed by Terraform"

  # Add additional configuration as needed
}
