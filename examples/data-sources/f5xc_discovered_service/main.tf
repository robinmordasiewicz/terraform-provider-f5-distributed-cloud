# Example configuration for f5xc_discovered_service data source

data "f5xc_discovered_service" "example" {
  name      = "existing-discovered_service"
  namespace = "system"
}

output "discovered_service_id" {
  value = data.f5xc_discovered_service.example.id
}
