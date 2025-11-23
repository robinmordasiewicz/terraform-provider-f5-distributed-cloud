# Example configuration for f5_distributed_cloud_allowed_tenant

resource "f5_distributed_cloud_allowed_tenant" "example" {
  name        = "example-allowed_tenant"
  namespace   = "system"
  description = "Example AllowedTenant resource managed by Terraform"

  # Add additional configuration as needed
}
