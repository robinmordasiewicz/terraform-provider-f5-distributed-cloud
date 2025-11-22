# Example configuration for f5xc_app_firewall

resource "f5xc_app_firewall" "example" {
  name        = "example-app_firewall"
  namespace   = "system"
  description = "Example AppFirewall resource managed by Terraform"

  # Add additional configuration as needed
}
