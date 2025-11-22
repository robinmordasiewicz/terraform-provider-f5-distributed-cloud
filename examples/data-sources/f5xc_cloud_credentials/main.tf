# Example configuration for f5xc_cloud_credentials data source

data "f5xc_cloud_credentials" "example" {
  name      = "existing-cloud_credentials"
  namespace = "system"
}

output "cloud_credentials_id" {
  value = data.f5xc_cloud_credentials.example.id
}
