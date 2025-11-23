# Example configuration for f5_distributed_cloud_log data source

data "f5_distributed_cloud_log" "example" {
  name      = "existing-log"
  namespace = "system"
}

output "log_id" {
  value = data.f5_distributed_cloud_log.example.id
}
