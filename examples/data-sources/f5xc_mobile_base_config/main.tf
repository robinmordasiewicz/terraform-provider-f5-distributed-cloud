# Example configuration for f5xc_mobile_base_config data source

data "f5xc_mobile_base_config" "example" {
  name      = "existing-mobile_base_config"
  namespace = "system"
}

output "mobile_base_config_id" {
  value = data.f5xc_mobile_base_config.example.id
}
