# Example configuration for f5_distributed_cloud_report_config data source

data "f5_distributed_cloud_report_config" "example" {
  name      = "existing-report_config"
  namespace = "system"
}

output "report_config_id" {
  value = data.f5_distributed_cloud_report_config.example.id
}
