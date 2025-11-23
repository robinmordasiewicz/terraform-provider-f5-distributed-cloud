# Example configuration for f5_distributed_cloud_receiver data source

data "f5_distributed_cloud_receiver" "example" {
  name      = "existing-receiver"
  namespace = "system"
}

output "receiver_id" {
  value = data.f5_distributed_cloud_receiver.example.id
}
