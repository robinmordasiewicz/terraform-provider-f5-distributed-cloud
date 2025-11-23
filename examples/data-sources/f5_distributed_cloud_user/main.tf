# Example configuration for f5_distributed_cloud_user data source

data "f5_distributed_cloud_user" "example" {
  name      = "existing-user"
  namespace = "system"
}

output "user_id" {
  value = data.f5_distributed_cloud_user.example.id
}
