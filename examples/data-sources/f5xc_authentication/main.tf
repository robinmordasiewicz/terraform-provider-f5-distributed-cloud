# Example configuration for f5xc_authentication data source

data "f5xc_authentication" "example" {
  name      = "existing-authentication"
  namespace = "system"
}

output "authentication_id" {
  value = data.f5xc_authentication.example.id
}
