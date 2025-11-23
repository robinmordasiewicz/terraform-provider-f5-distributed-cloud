# Example configuration for f5_distributed_cloud_mitigated_domain

resource "f5_distributed_cloud_mitigated_domain" "example" {
  name        = "example-mitigated_domain"
  namespace   = "system"
  description = "Example MitigatedDomain resource managed by Terraform"

  # Add additional configuration as needed
}
