# Example configuration for f5_distributed_cloud_nginx_service_discovery

resource "f5_distributed_cloud_nginx_service_discovery" "example" {
  name        = "example-nginx_service_discovery"
  namespace   = "system"
  description = "Example NginxServiceDiscovery resource managed by Terraform"

  # Add additional configuration as needed
}
