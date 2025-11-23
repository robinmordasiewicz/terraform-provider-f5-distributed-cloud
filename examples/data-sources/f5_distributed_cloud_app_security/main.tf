# Example configuration for f5_distributed_cloud_app_security data source

data "f5_distributed_cloud_app_security" "example" {
  name      = "existing-app_security"
  namespace = "system"
}

output "app_security_id" {
  value = data.f5_distributed_cloud_app_security.example.id
}
