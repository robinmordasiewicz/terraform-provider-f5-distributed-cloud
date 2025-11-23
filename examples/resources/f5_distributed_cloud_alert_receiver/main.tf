# Example configuration for f5_distributed_cloud_alert_receiver

resource "f5_distributed_cloud_alert_receiver" "example" {
  name        = "example-alert_receiver"
  namespace   = "system"
  description = "Example AlertReceiver resource managed by Terraform"

  # Add additional configuration as needed
}
