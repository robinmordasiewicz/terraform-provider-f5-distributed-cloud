# Example configuration for f5_distributed_cloud_mitigated_domain data source

data "f5_distributed_cloud_mitigated_domain" "example" {
  name      = "existing-mitigated_domain"
  namespace = "system"
}

output "mitigated_domain_id" {
  value = data.f5_distributed_cloud_mitigated_domain.example.id
}
