# Example configuration for f5xc_segment_connection data source

data "f5xc_segment_connection" "example" {
  name      = "existing-segment_connection"
  namespace = "system"
}

output "segment_connection_id" {
  value = data.f5xc_segment_connection.example.id
}
