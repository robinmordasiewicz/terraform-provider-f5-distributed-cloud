# Example configuration for f5_distributed_cloud_app_type data source

data "f5_distributed_cloud_app_type" "example" {
  name      = "existing-app_type"
  namespace = "system"
}

output "app_type_id" {
  value = data.f5_distributed_cloud_app_type.example.id
}
