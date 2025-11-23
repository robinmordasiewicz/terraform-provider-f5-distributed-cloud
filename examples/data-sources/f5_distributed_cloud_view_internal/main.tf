# Example configuration for f5_distributed_cloud_view_internal data source

data "f5_distributed_cloud_view_internal" "example" {
  name      = "existing-view_internal"
  namespace = "system"
}

output "view_internal_id" {
  value = data.f5_distributed_cloud_view_internal.example.id
}
