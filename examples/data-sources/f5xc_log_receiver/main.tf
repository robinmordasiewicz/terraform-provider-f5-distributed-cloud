# Example configuration for f5xc_log_receiver data source

data "f5xc_log_receiver" "example" {
  name      = "existing-log_receiver"
  namespace = "system"
}

output "log_receiver_id" {
  value = data.f5xc_log_receiver.example.id
}
