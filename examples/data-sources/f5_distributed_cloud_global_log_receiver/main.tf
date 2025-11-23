# Example configuration for f5_distributed_cloud_global_log_receiver data source

data "f5_distributed_cloud_global_log_receiver" "example" {
  name      = "existing-global_log_receiver"
  namespace = "system"
}

output "global_log_receiver_id" {
  value = data.f5_distributed_cloud_global_log_receiver.example.id
}
