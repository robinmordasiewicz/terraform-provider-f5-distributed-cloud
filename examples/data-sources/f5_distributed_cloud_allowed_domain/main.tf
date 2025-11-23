# Example configuration for f5_distributed_cloud_allowed_domain data source

data "f5_distributed_cloud_allowed_domain" "example" {
  name      = "existing-allowed_domain"
  namespace = "system"
}

output "allowed_domain_id" {
  value = data.f5_distributed_cloud_allowed_domain.example.id
}
