# Example configuration for f5xc_cloud_link data source

data "f5xc_cloud_link" "example" {
  name      = "existing-cloud_link"
  namespace = "system"
}

output "cloud_link_id" {
  value = data.f5xc_cloud_link.example.id
}
