# Example configuration for f5_distributed_cloud_secret_management_access data source

data "f5_distributed_cloud_secret_management_access" "example" {
  name      = "existing-secret_management_access"
  namespace = "system"
}

output "secret_management_access_id" {
  value = data.f5_distributed_cloud_secret_management_access.example.id
}
