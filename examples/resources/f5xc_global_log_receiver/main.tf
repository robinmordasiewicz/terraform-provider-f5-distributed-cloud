# Example configuration for f5xc_global_log_receiver

resource "f5xc_global_log_receiver" "example" {
  name        = "example-global_log_receiver"
  namespace   = "system"
  description = "Example GlobalLogReceiver resource managed by Terraform"

  # Add additional configuration as needed
}
