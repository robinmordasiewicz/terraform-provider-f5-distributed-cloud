# Example configuration for f5xc_protocol_inspection data source

data "f5xc_protocol_inspection" "example" {
  name      = "existing-protocol_inspection"
  namespace = "system"
}

output "protocol_inspection_id" {
  value = data.f5xc_protocol_inspection.example.id
}
