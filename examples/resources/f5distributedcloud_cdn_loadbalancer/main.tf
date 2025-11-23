# Example configuration for f5distributedcloud_cdn_loadbalancer

resource "f5distributedcloud_cdn_loadbalancer" "example" {
  name        = "example-cdn_loadbalancer"
  namespace   = "system"
  description = "Example CDNLoadbalancer resource managed by Terraform"

  # Add additional configuration as needed
}
