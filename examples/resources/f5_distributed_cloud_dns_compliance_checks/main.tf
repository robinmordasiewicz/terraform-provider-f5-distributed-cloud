# Example configuration for f5_distributed_cloud_dns_compliance_checks

resource "f5_distributed_cloud_dns_compliance_checks" "example" {
  name        = "example-dns_compliance_checks"
  namespace   = "system"
  description = "Example DNSComplianceChecks resource managed by Terraform"

  # Add additional configuration as needed
}
