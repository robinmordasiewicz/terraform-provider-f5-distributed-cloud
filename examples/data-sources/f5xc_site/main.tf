# Example configuration for f5xc_site data source

data "f5xc_site" "example" {
  name      = "existing-site"
  namespace = "system"
}

output "site_id" {
  value = data.f5xc_site.example.id
}
