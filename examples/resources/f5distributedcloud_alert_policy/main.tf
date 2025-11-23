# Example configuration for f5distributedcloud_alert_policy

resource "f5distributedcloud_alert_policy" "example" {
  name        = "example-alert_policy"
  namespace   = "system"
  description = "Example AlertPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
