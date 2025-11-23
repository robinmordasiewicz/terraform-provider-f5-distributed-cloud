# Example configuration for f5_distributed_cloud_static_component data source

data "f5_distributed_cloud_static_component" "example" {
  name      = "existing-static_component"
  namespace = "system"
}

output "static_component_id" {
  value = data.f5_distributed_cloud_static_component.example.id
}
