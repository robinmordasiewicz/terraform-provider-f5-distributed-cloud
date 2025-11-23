# Example configuration for f5_distributed_cloud_alert_template

resource "f5_distributed_cloud_alert_template" "example" {
  name        = "example-alert_template"
  namespace   = "system"
  description = "Example AlertTemplate resource managed by Terraform"

  # Add additional configuration as needed
}
