# Example configuration for f5xc_dns_domain data source

data "f5xc_dns_domain" "example" {
  name      = "existing-dns_domain"
  namespace = "system"
}

output "dns_domain_id" {
  value = data.f5xc_dns_domain.example.id
}
