# Example configuration for f5distributedcloud_site data source

data "f5distributedcloud_site" "example" {
  name      = "existing-site"
  namespace = "system"
}

output "site_id" {
  value = data.f5distributedcloud_site.example.id
}
