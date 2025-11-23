# Example configuration for f5_distributed_cloud_filter_set data source

data "f5_distributed_cloud_filter_set" "example" {
  name      = "existing-filter_set"
  namespace = "system"
}

output "filter_set_id" {
  value = data.f5_distributed_cloud_filter_set.example.id
}
