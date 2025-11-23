# Example configuration for f5_distributed_cloud_endpoint data source

data "f5_distributed_cloud_endpoint" "example" {
  name      = "existing-endpoint"
  namespace = "system"
}

output "endpoint_id" {
  value = data.f5_distributed_cloud_endpoint.example.id
}
