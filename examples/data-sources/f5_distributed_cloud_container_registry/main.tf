# Example configuration for f5_distributed_cloud_container_registry data source

data "f5_distributed_cloud_container_registry" "example" {
  name      = "existing-container_registry"
  namespace = "system"
}

output "container_registry_id" {
  value = data.f5_distributed_cloud_container_registry.example.id
}
