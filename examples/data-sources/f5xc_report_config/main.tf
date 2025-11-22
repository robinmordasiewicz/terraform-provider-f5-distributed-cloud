# Example configuration for f5xc_report_config data source

data "f5xc_report_config" "example" {
  name      = "existing-report_config"
  namespace = "system"
}

output "report_config_id" {
  value = data.f5xc_report_config.example.id
}
