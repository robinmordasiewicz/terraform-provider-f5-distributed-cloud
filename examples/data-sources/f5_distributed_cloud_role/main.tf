# Example configuration for f5_distributed_cloud_role data source

data "f5_distributed_cloud_role" "example" {
  name      = "existing-role"
  namespace = "system"
}

output "role_id" {
  value = data.f5_distributed_cloud_role.example.id
}
