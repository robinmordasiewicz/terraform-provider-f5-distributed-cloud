# Example configuration for f5_distributed_cloud_cloud_connect data source

data "f5_distributed_cloud_cloud_connect" "example" {
  name      = "existing-cloud_connect"
  namespace = "system"
}

output "cloud_connect_id" {
  value = data.f5_distributed_cloud_cloud_connect.example.id
}
