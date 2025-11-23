# Example configuration for f5distributedcloud_role data source

data "f5distributedcloud_role" "example" {
  name      = "existing-role"
  namespace = "system"
}

output "role_id" {
  value = data.f5distributedcloud_role.example.id
}
