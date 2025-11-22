# Example configuration for f5xc_managed_tenant

resource "f5xc_managed_tenant" "example" {
  name        = "example-managed_tenant"
  namespace   = "system"
  description = "Example ManagedTenant resource managed by Terraform"

  # Add additional configuration as needed
}
