# Example configuration for f5_distributed_cloud_app_setting data source

data "f5_distributed_cloud_app_setting" "example" {
  name      = "existing-app_setting"
  namespace = "system"
}

output "app_setting_id" {
  value = data.f5_distributed_cloud_app_setting.example.id
}
