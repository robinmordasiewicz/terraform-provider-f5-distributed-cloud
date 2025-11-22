# Example configuration for f5xc_allowed_domain

resource "f5xc_allowed_domain" "example" {
  name        = "example-allowed_domain"
  namespace   = "system"
  description = "Example AllowedDomain resource managed by Terraform"

  # Add additional configuration as needed
}
