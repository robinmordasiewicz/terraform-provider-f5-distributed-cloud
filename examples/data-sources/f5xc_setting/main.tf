# Example configuration for f5xc_setting data source

data "f5xc_setting" "example" {
  name      = "existing-setting"
  namespace = "system"
}

output "setting_id" {
  value = data.f5xc_setting.example.id
}
