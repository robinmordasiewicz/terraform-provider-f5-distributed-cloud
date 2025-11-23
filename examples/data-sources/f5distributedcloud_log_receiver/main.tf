# Example configuration for f5distributedcloud_log_receiver data source

data "f5distributedcloud_log_receiver" "example" {
  name      = "existing-log_receiver"
  namespace = "system"
}

output "log_receiver_id" {
  value = data.f5distributedcloud_log_receiver.example.id
}
