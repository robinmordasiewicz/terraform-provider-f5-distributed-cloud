# Example configuration for f5_distributed_cloud_registration data source

data "f5_distributed_cloud_registration" "example" {
  name      = "existing-registration"
  namespace = "system"
}

output "registration_id" {
  value = data.f5_distributed_cloud_registration.example.id
}
