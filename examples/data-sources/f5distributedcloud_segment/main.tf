# Example configuration for f5distributedcloud_segment data source

data "f5distributedcloud_segment" "example" {
  name      = "existing-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.f5distributedcloud_segment.example.id
}
