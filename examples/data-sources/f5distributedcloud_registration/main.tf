# Example configuration for f5distributedcloud_registration data source

data "f5distributedcloud_registration" "example" {
  name      = "existing-registration"
  namespace = "system"
}

output "registration_id" {
  value = data.f5distributedcloud_registration.example.id
}
