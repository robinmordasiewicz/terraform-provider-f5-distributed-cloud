# Example configuration for f5xc_proxy

resource "f5xc_proxy" "example" {
  name        = "example-proxy"
  namespace   = "system"
  description = "Example Proxy resource managed by Terraform"

  # Add additional configuration as needed
}
