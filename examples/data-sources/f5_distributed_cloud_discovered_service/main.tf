# Example configuration for f5_distributed_cloud_discovered_service data source

data "f5_distributed_cloud_discovered_service" "example" {
  name      = "existing-discovered_service"
  namespace = "system"
}

output "discovered_service_id" {
  value = data.f5_distributed_cloud_discovered_service.example.id
}
