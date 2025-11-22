# Example: Look up an existing app type
data "f5xc_app_type" "example" {
  name      = "my-app-type"
  namespace = "shared"
}

# Output the app type
output "app_type" {
  value = data.f5xc_app_type.example.app_type
}

output "app_type_id" {
  value = data.f5xc_app_type.example.id
}
