# Example configuration for f5xc_tenant_configuration data source

data "f5xc_tenant_configuration" "example" {
  name      = "existing-tenant_configuration"
  namespace = "system"
}

output "tenant_configuration_id" {
  value = data.f5xc_tenant_configuration.example.id
}
