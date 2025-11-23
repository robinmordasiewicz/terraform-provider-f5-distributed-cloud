# Example configuration for f5_distributed_cloud_secret_management_access

resource "f5_distributed_cloud_secret_management_access" "example" {
  name        = "example-secret_management_access"
  namespace   = "system"
  description = "Example SecretManagementAccess resource managed by Terraform"

  # Add additional configuration as needed
}
