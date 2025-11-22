# Example configuration for f5xc_cloud_region data source

data "f5xc_cloud_region" "example" {
  name      = "existing-cloud_region"
  namespace = "system"
}

output "cloud_region_id" {
  value = data.f5xc_cloud_region.example.id
}
