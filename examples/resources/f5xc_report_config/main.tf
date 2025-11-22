# Example configuration for f5xc_report_config

resource "f5xc_report_config" "example" {
  name        = "example-report_config"
  namespace   = "system"
  description = "Example ReportConfig resource managed by Terraform"

  # Add additional configuration as needed
}
