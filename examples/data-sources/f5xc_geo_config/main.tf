# Example configuration for f5xc_geo_config data source

data "f5xc_geo_config" "example" {
  name      = "existing-geo_config"
  namespace = "system"
}

output "geo_config_id" {
  value = data.f5xc_geo_config.example.id
}
