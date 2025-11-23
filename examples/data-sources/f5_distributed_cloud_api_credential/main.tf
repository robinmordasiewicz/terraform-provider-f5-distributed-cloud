# Example configuration for f5_distributed_cloud_api_credential data source

data "f5_distributed_cloud_api_credential" "example" {
  name      = "existing-api_credential"
  namespace = "system"
}

output "api_credential_id" {
  value = data.f5_distributed_cloud_api_credential.example.id
}
