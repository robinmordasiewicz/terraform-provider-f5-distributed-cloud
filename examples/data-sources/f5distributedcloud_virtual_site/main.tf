# Example configuration for f5distributedcloud_virtual_site data source

data "f5distributedcloud_virtual_site" "example" {
  name      = "existing-virtual_site"
  namespace = "system"
}

output "virtual_site_id" {
  value = data.f5distributedcloud_virtual_site.example.id
}
