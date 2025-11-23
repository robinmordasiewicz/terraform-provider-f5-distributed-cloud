# Example configuration for f5_distributed_cloud_protected_application

resource "f5_distributed_cloud_protected_application" "example" {
  name        = "example-protected_application"
  namespace   = "system"
  description = "Example ProtectedApplication resource managed by Terraform"

  # Add additional configuration as needed
}
