# Example configuration for f5xc_namespace_role data source

data "f5xc_namespace_role" "example" {
  name      = "existing-namespace_role"
  namespace = "system"
}

output "namespace_role_id" {
  value = data.f5xc_namespace_role.example.id
}
