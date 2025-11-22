# Example configuration for f5xc_user data source

data "f5xc_user" "example" {
  name      = "existing-user"
  namespace = "system"
}

output "user_id" {
  value = data.f5xc_user.example.id
}
