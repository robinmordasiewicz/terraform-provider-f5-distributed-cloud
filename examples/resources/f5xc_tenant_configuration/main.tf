# Example configuration for f5xc_tenant_configuration

resource "f5xc_tenant_configuration" "example" {
  name        = "example-tenant_configuration"
  namespace   = "system"
  description = "Example TenantConfiguration resource managed by Terraform"

  # Add additional configuration as needed
}
