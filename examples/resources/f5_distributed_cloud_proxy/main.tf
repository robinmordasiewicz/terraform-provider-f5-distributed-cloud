# Example configuration for f5_distributed_cloud_proxy

resource "f5_distributed_cloud_proxy" "example" {
  name        = "example-proxy"
  namespace   = "system"
  description = "Example Proxy resource managed by Terraform"

  # Add additional configuration as needed
}
