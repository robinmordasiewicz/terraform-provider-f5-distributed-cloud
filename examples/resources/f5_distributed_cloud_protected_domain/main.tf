# Example configuration for f5_distributed_cloud_protected_domain

resource "f5_distributed_cloud_protected_domain" "example" {
  name        = "example-protected_domain"
  namespace   = "system"
  description = "Example ProtectedDomain resource managed by Terraform"

  # Add additional configuration as needed
}
