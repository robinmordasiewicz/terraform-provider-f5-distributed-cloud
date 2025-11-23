# Example configuration for f5_distributed_cloud_user_group data source

data "f5_distributed_cloud_user_group" "example" {
  name      = "existing-user_group"
  namespace = "system"
}

output "user_group_id" {
  value = data.f5_distributed_cloud_user_group.example.id
}
