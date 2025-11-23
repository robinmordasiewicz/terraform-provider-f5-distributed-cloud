# Example configuration for f5_distributed_cloud_cdn_loadbalancer

resource "f5_distributed_cloud_cdn_loadbalancer" "example" {
  name        = "example-cdn_loadbalancer"
  namespace   = "system"
  description = "Example CDNLoadbalancer resource managed by Terraform"

  # Add additional configuration as needed
}
