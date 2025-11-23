# Example configuration for f5_distributed_cloud_nginx_service_discovery data source

data "f5_distributed_cloud_nginx_service_discovery" "example" {
  name      = "existing-nginx_service_discovery"
  namespace = "system"
}

output "nginx_service_discovery_id" {
  value = data.f5_distributed_cloud_nginx_service_discovery.example.id
}
