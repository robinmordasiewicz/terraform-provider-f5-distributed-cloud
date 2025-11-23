# Example configuration for f5_distributed_cloud_alert_policy

resource "f5_distributed_cloud_alert_policy" "example" {
  name        = "example-alert_policy"
  namespace   = "system"
  description = "Example AlertPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
