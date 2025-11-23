# Example configuration for f5_distributed_cloud_debug data source

data "f5_distributed_cloud_debug" "example" {
  name      = "existing-debug"
  namespace = "system"
}

output "debug_id" {
  value = data.f5_distributed_cloud_debug.example.id
}
