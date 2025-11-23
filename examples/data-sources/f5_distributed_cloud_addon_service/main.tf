# Example configuration for f5_distributed_cloud_addon_service data source

data "f5_distributed_cloud_addon_service" "example" {
  name      = "existing-addon_service"
  namespace = "system"
}

output "addon_service_id" {
  value = data.f5_distributed_cloud_addon_service.example.id
}
