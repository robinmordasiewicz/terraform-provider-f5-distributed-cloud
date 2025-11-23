# Example configuration for f5distributedcloud_global_log_receiver data source

data "f5distributedcloud_global_log_receiver" "example" {
  name      = "existing-global_log_receiver"
  namespace = "system"
}

output "global_log_receiver_id" {
  value = data.f5distributedcloud_global_log_receiver.example.id
}
