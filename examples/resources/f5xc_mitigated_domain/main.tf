# Example configuration for f5xc_mitigated_domain

resource "f5xc_mitigated_domain" "example" {
  name        = "example-mitigated_domain"
  namespace   = "system"
  description = "Example MitigatedDomain resource managed by Terraform"

  # Add additional configuration as needed
}
