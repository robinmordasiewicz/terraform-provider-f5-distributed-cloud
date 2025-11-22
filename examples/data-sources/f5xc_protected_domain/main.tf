# Example configuration for f5xc_protected_domain data source

data "f5xc_protected_domain" "example" {
  name      = "existing-protected_domain"
  namespace = "system"
}

output "protected_domain_id" {
  value = data.f5xc_protected_domain.example.id
}
