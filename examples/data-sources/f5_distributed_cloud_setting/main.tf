# Example configuration for f5_distributed_cloud_setting data source

data "f5_distributed_cloud_setting" "example" {
  name      = "existing-setting"
  namespace = "system"
}

output "setting_id" {
  value = data.f5_distributed_cloud_setting.example.id
}
