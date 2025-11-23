# Example configuration for f5_distributed_cloud_segment data source

data "f5_distributed_cloud_segment" "example" {
  name      = "existing-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.f5_distributed_cloud_segment.example.id
}
