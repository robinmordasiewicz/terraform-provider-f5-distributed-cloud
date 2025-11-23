# Example configuration for f5_distributed_cloud_virtual_site data source

data "f5_distributed_cloud_virtual_site" "example" {
  name      = "existing-virtual_site"
  namespace = "system"
}

output "virtual_site_id" {
  value = data.f5_distributed_cloud_virtual_site.example.id
}
