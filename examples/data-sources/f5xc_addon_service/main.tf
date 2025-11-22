# Example configuration for f5xc_addon_service data source

data "f5xc_addon_service" "example" {
  name      = "existing-addon_service"
  namespace = "system"
}

output "addon_service_id" {
  value = data.f5xc_addon_service.example.id
}
