# Example configuration for f5xc_protected_domain

resource "f5xc_protected_domain" "example" {
  name        = "example-protected_domain"
  namespace   = "system"
  description = "Example ProtectedDomain resource managed by Terraform"

  # Add additional configuration as needed
}
