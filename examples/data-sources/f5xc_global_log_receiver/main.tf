# Example: Look up an existing global log receiver
data "f5xc_global_log_receiver" "example" {
  name      = "my-global-log-receiver"
  namespace = "system"
}

# Output the receiver type
output "receiver_type" {
  value = data.f5xc_global_log_receiver.example.receiver_type
}

output "global_log_receiver_id" {
  value = data.f5xc_global_log_receiver.example.id
}
