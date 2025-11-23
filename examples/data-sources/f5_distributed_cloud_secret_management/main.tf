# Example configuration for f5_distributed_cloud_secret_management data source

data "f5_distributed_cloud_secret_management" "example" {
  name      = "existing-secret_management"
  namespace = "system"
}

output "secret_management_id" {
  value = data.f5_distributed_cloud_secret_management.example.id
}
