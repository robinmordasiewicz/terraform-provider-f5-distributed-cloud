# Example configuration for f5xc_registration data source

data "f5xc_registration" "example" {
  name      = "existing-registration"
  namespace = "system"
}

output "registration_id" {
  value = data.f5xc_registration.example.id
}
