# Example configuration for f5xc_nginx_service_discovery

resource "f5xc_nginx_service_discovery" "example" {
  name        = "example-nginx_service_discovery"
  namespace   = "system"
  description = "Example NginxServiceDiscovery resource managed by Terraform"

  # Add additional configuration as needed
}
