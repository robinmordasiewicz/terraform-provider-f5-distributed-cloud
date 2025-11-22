# Example configuration for f5xc_segment data source

data "f5xc_segment" "example" {
  name      = "existing-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.f5xc_segment.example.id
}
