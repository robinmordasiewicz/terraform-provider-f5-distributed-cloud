# Example configuration for f5_distributed_cloud_v1_http_monitor data source

data "f5_distributed_cloud_v1_http_monitor" "example" {
  name      = "existing-v1_http_monitor"
  namespace = "system"
}

output "v1_http_monitor_id" {
  value = data.f5_distributed_cloud_v1_http_monitor.example.id
}
