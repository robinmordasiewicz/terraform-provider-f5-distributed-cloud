# Example configuration for f5xc_voltstack_site data source

data "f5xc_voltstack_site" "example" {
  name      = "existing-voltstack_site"
  namespace = "system"
}

output "voltstack_site_id" {
  value = data.f5xc_voltstack_site.example.id
}
