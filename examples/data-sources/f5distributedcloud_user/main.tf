# Example configuration for f5distributedcloud_user data source

data "f5distributedcloud_user" "example" {
  name      = "existing-user"
  namespace = "system"
}

output "user_id" {
  value = data.f5distributedcloud_user.example.id
}
