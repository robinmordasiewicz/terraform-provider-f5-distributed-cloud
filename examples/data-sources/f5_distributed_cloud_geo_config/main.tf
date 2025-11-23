# Example configuration for f5_distributed_cloud_geo_config data source

data "f5_distributed_cloud_geo_config" "example" {
  name      = "existing-geo_config"
  namespace = "system"
}

output "geo_config_id" {
  value = data.f5_distributed_cloud_geo_config.example.id
}
