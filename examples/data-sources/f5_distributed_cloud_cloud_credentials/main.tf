# Example configuration for f5_distributed_cloud_cloud_credentials data source

data "f5_distributed_cloud_cloud_credentials" "example" {
  name      = "existing-cloud_credentials"
  namespace = "system"
}

output "cloud_credentials_id" {
  value = data.f5_distributed_cloud_cloud_credentials.example.id
}
