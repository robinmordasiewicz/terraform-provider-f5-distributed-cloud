# Example configuration for f5_distributed_cloud_tenant_configuration data source

data "f5_distributed_cloud_tenant_configuration" "example" {
  name      = "existing-tenant_configuration"
  namespace = "system"
}

output "tenant_configuration_id" {
  value = data.f5_distributed_cloud_tenant_configuration.example.id
}
