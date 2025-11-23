# Example configuration for f5_distributed_cloud_role

resource "f5_distributed_cloud_role" "example" {
  name        = "example-role"
  namespace   = "system"
  description = "Example Role resource managed by Terraform"

  # Add additional configuration as needed
}
