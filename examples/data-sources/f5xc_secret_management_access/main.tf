# Example configuration for f5xc_secret_management_access data source

data "f5xc_secret_management_access" "example" {
  name      = "existing-secret_management_access"
  namespace = "system"
}

output "secret_management_access_id" {
  value = data.f5xc_secret_management_access.example.id
}
