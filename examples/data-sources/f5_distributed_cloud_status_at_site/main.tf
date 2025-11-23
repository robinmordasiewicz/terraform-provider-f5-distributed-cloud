# Example configuration for f5_distributed_cloud_status_at_site data source

data "f5_distributed_cloud_status_at_site" "example" {
  name      = "existing-status_at_site"
  namespace = "system"
}

output "status_at_site_id" {
  value = data.f5_distributed_cloud_status_at_site.example.id
}
