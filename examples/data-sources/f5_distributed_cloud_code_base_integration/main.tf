# Example configuration for f5_distributed_cloud_code_base_integration data source

data "f5_distributed_cloud_code_base_integration" "example" {
  name      = "existing-code_base_integration"
  namespace = "system"
}

output "code_base_integration_id" {
  value = data.f5_distributed_cloud_code_base_integration.example.id
}
