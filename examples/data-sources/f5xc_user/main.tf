# Example: Look up an existing user
data "f5xc_user" "example" {
  name      = "my-user"
  namespace = "system"
}

# Output the user email
output "email" {
  value = data.f5xc_user.example.email
}

output "user_id" {
  value = data.f5xc_user.example.id
}
