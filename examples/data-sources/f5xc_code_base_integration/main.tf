# Example configuration for f5xc_code_base_integration data source

data "f5xc_code_base_integration" "example" {
  name      = "existing-code_base_integration"
  namespace = "system"
}

output "code_base_integration_id" {
  value = data.f5xc_code_base_integration.example.id
}
