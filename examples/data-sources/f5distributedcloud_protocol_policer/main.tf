# Example configuration for f5distributedcloud_protocol_policer data source

data "f5distributedcloud_protocol_policer" "example" {
  name      = "existing-protocol_policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.f5distributedcloud_protocol_policer.example.id
}
