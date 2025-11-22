# Example: Look up an existing role
data "f5xc_role" "example" {
  name      = "my-role"
  namespace = "system"
}

# Output the role type
output "role_type" {
  value = data.f5xc_role.example.role_type
}

output "role_id" {
  value = data.f5xc_role.example.id
}
