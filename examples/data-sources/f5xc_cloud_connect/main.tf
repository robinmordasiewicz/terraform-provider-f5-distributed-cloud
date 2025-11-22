# Example configuration for f5xc_cloud_connect data source

data "f5xc_cloud_connect" "example" {
  name      = "existing-cloud_connect"
  namespace = "system"
}

output "cloud_connect_id" {
  value = data.f5xc_cloud_connect.example.id
}
