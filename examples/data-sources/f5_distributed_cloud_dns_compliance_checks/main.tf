# Example configuration for f5_distributed_cloud_dns_compliance_checks data source

data "f5_distributed_cloud_dns_compliance_checks" "example" {
  name      = "existing-dns_compliance_checks"
  namespace = "system"
}

output "dns_compliance_checks_id" {
  value = data.f5_distributed_cloud_dns_compliance_checks.example.id
}
