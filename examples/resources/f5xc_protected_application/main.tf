# Example configuration for f5xc_protected_application

resource "f5xc_protected_application" "example" {
  name        = "example-protected_application"
  namespace   = "system"
  description = "Example ProtectedApplication resource managed by Terraform"

  # Add additional configuration as needed
}
