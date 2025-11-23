# Example configuration for f5_distributed_cloud_dns_domain data source

data "f5_distributed_cloud_dns_domain" "example" {
  name      = "existing-dns_domain"
  namespace = "system"
}

output "dns_domain_id" {
  value = data.f5_distributed_cloud_dns_domain.example.id
}
