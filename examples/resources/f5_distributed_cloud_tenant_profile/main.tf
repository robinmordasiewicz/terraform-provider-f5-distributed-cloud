# Example configuration for f5_distributed_cloud_tenant_profile

resource "f5_distributed_cloud_tenant_profile" "example" {
  name        = "example-tenant_profile"
  namespace   = "system"
  description = "Example TenantProfile resource managed by Terraform"

  # Add additional configuration as needed
}
