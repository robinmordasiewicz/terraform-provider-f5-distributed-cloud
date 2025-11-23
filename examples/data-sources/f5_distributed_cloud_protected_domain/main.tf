# Example configuration for f5_distributed_cloud_protected_domain data source

data "f5_distributed_cloud_protected_domain" "example" {
  name      = "existing-protected_domain"
  namespace = "system"
}

output "protected_domain_id" {
  value = data.f5_distributed_cloud_protected_domain.example.id
}
