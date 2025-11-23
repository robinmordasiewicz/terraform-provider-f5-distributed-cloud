# Example configuration for f5_distributed_cloud_protocol_policer data source

data "f5_distributed_cloud_protocol_policer" "example" {
  name      = "existing-protocol_policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.f5_distributed_cloud_protocol_policer.example.id
}
