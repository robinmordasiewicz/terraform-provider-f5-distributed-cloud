# Example configuration for f5_distributed_cloud_data_group data source

data "f5_distributed_cloud_data_group" "example" {
  name      = "existing-data_group"
  namespace = "system"
}

output "data_group_id" {
  value = data.f5_distributed_cloud_data_group.example.id
}
