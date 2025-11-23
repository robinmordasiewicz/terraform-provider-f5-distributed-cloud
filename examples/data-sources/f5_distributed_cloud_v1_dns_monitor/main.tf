# Example configuration for f5_distributed_cloud_v1_dns_monitor data source

data "f5_distributed_cloud_v1_dns_monitor" "example" {
  name      = "existing-v1_dns_monitor"
  namespace = "system"
}

output "v1_dns_monitor_id" {
  value = data.f5_distributed_cloud_v1_dns_monitor.example.id
}
