# Example configuration for f5_distributed_cloud_protected_application data source

data "f5_distributed_cloud_protected_application" "example" {
  name      = "existing-protected_application"
  namespace = "system"
}

output "protected_application_id" {
  value = data.f5_distributed_cloud_protected_application.example.id
}
