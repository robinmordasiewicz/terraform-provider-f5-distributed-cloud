# Example configuration for f5xc_alert_gen_policy

resource "f5xc_alert_gen_policy" "example" {
  name        = "example-alert_gen_policy"
  namespace   = "system"
  description = "Example AlertGenPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
