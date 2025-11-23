# Example configuration for f5_distributed_cloud_authentication data source

data "f5_distributed_cloud_authentication" "example" {
  name      = "existing-authentication"
  namespace = "system"
}

output "authentication_id" {
  value = data.f5_distributed_cloud_authentication.example.id
}
