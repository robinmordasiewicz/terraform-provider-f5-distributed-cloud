# Example configuration for f5_distributed_cloud_child_tenant_manager

resource "f5_distributed_cloud_child_tenant_manager" "example" {
  name        = "example-child_tenant_manager"
  namespace   = "system"
  description = "Example ChildTenantManager resource managed by Terraform"

  # Add additional configuration as needed
}
