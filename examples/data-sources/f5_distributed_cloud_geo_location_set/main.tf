# Example configuration for f5_distributed_cloud_geo_location_set data source

data "f5_distributed_cloud_geo_location_set" "example" {
  name      = "existing-geo_location_set"
  namespace = "system"
}

output "geo_location_set_id" {
  value = data.f5_distributed_cloud_geo_location_set.example.id
}
