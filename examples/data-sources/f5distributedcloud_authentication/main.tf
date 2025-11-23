# Example configuration for f5distributedcloud_authentication data source

data "f5distributedcloud_authentication" "example" {
  name      = "existing-authentication"
  namespace = "system"
}

output "authentication_id" {
  value = data.f5distributedcloud_authentication.example.id
}
