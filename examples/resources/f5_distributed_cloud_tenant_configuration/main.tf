# Example configuration for f5_distributed_cloud_tenant_configuration

resource "f5_distributed_cloud_tenant_configuration" "example" {
  name        = "example-tenant_configuration"
  namespace   = "system"
  description = "Example TenantConfiguration resource managed by Terraform"

  # Add additional configuration as needed
}
