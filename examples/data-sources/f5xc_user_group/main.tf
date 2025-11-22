# Example configuration for f5xc_user_group data source

data "f5xc_user_group" "example" {
  name      = "existing-user_group"
  namespace = "system"
}

output "user_group_id" {
  value = data.f5xc_user_group.example.id
}
