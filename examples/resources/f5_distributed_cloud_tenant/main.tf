# Example configuration for f5_distributed_cloud_tenant

resource "f5_distributed_cloud_tenant" "example" {
  name        = "example-tenant"
  namespace   = "system"
  description = "Example Tenant resource managed by Terraform"

  # Add additional configuration as needed
}
