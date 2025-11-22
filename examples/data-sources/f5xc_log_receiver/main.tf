# Example: Look up an existing log receiver
data "f5xc_log_receiver" "example" {
  name      = "my-log-receiver"
  namespace = "my-namespace"
}

# Output the receiver type
output "receiver_type" {
  value = data.f5xc_log_receiver.example.receiver_type
}

output "log_receiver_id" {
  value = data.f5xc_log_receiver.example.id
}
