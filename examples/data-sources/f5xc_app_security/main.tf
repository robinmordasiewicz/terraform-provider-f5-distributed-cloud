# Example configuration for f5xc_app_security data source

data "f5xc_app_security" "example" {
  name      = "existing-app_security"
  namespace = "system"
}

output "app_security_id" {
  value = data.f5xc_app_security.example.id
}
