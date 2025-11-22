# Example configuration for f5xc_static_component data source

data "f5xc_static_component" "example" {
  name      = "existing-static_component"
  namespace = "system"
}

output "static_component_id" {
  value = data.f5xc_static_component.example.id
}
