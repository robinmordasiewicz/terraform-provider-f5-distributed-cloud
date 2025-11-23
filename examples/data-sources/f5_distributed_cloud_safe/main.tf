# Example configuration for f5_distributed_cloud_safe data source

data "f5_distributed_cloud_safe" "example" {
  name      = "existing-safe"
  namespace = "system"
}

output "safe_id" {
  value = data.f5_distributed_cloud_safe.example.id
}
