# Example configuration for f5xc_allowed_tenant

resource "f5xc_allowed_tenant" "example" {
  name        = "example-allowed_tenant"
  namespace   = "system"
  description = "Example AllowedTenant resource managed by Terraform"

  # Add additional configuration as needed
}
