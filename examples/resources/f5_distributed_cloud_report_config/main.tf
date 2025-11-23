# Example configuration for f5_distributed_cloud_report_config

resource "f5_distributed_cloud_report_config" "example" {
  name        = "example-report_config"
  namespace   = "system"
  description = "Example ReportConfig resource managed by Terraform"

  # Add additional configuration as needed
}
