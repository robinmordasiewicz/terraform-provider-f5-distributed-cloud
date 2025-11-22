# Example configuration for f5xc_receiver data source

data "f5xc_receiver" "example" {
  name      = "existing-receiver"
  namespace = "system"
}

output "receiver_id" {
  value = data.f5xc_receiver.example.id
}
