# Example configuration for f5distributedcloud_log_receiver

resource "f5distributedcloud_log_receiver" "example" {
  name        = "example-log_receiver"
  namespace   = "system"
  description = "Example LogReceiver resource managed by Terraform"

  # Add additional configuration as needed
}
