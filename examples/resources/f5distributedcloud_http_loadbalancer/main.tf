# Example configuration for f5distributedcloud_http_loadbalancer

resource "f5distributedcloud_http_loadbalancer" "example" {
  name        = "example-http_loadbalancer"
  namespace   = "system"
  description = "Example HTTPLoadbalancer resource managed by Terraform"

  # Add additional configuration as needed
}
