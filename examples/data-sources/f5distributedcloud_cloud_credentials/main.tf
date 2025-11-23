# Example configuration for f5distributedcloud_cloud_credentials data source

data "f5distributedcloud_cloud_credentials" "example" {
  name      = "existing-cloud_credentials"
  namespace = "system"
}

output "cloud_credentials_id" {
  value = data.f5distributedcloud_cloud_credentials.example.id
}
