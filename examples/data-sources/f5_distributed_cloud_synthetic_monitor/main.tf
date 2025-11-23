# Example configuration for f5_distributed_cloud_synthetic_monitor data source

data "f5_distributed_cloud_synthetic_monitor" "example" {
  name      = "existing-synthetic_monitor"
  namespace = "system"
}

output "synthetic_monitor_id" {
  value = data.f5_distributed_cloud_synthetic_monitor.example.id
}
