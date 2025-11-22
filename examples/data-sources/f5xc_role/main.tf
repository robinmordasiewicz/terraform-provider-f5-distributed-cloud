# Example configuration for f5xc_role data source

data "f5xc_role" "example" {
  name      = "existing-role"
  namespace = "system"
}

output "role_id" {
  value = data.f5xc_role.example.id
}
