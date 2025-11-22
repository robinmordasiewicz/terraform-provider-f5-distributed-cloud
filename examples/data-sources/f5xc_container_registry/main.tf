# Example configuration for f5xc_container_registry data source

data "f5xc_container_registry" "example" {
  name      = "existing-container_registry"
  namespace = "system"
}

output "container_registry_id" {
  value = data.f5xc_container_registry.example.id
}
