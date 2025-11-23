# Example configuration for f5_distributed_cloud_site data source

data "f5_distributed_cloud_site" "example" {
  name      = "existing-site"
  namespace = "system"
}

output "site_id" {
  value = data.f5_distributed_cloud_site.example.id
}
