# Example configuration for f5xc_secret_management_access

resource "f5xc_secret_management_access" "example" {
  name        = "example-secret_management_access"
  namespace   = "system"
  description = "Example SecretManagementAccess resource managed by Terraform"

  # Add additional configuration as needed
}
