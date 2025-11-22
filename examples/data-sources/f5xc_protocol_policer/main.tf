# Example: Look up an existing protocol policer
data "f5xc_protocol_policer" "example" {
  name      = "my-protocol-policer"
  namespace = "my-namespace"
}

# Output the enabled status
output "enabled" {
  value = data.f5xc_protocol_policer.example.enabled
}

output "protocol_policer_id" {
  value = data.f5xc_protocol_policer.example.id
}
