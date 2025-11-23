# Example configuration for f5_distributed_cloud_segment_connection data source

data "f5_distributed_cloud_segment_connection" "example" {
  name      = "existing-segment_connection"
  namespace = "system"
}

output "segment_connection_id" {
  value = data.f5_distributed_cloud_segment_connection.example.id
}
