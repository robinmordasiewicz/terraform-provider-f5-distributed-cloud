# Example configuration for f5_distributed_cloud_child_tenant

resource "f5_distributed_cloud_child_tenant" "example" {
  name        = "example-child_tenant"
  namespace   = "system"
  description = "Example ChildTenant resource managed by Terraform"

  # Add additional configuration as needed
}
