# Example configuration for f5xc_tenant

resource "f5xc_tenant" "example" {
  name        = "example-tenant"
  namespace   = "system"
  description = "Example Tenant resource managed by Terraform"

  # Add additional configuration as needed
}
