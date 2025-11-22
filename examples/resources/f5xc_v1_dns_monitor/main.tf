# Example configuration for f5xc_v1_dns_monitor

resource "f5xc_v1_dns_monitor" "example" {
  name        = "example-v1_dns_monitor"
  namespace   = "system"
  description = "Example V1DNSMonitor resource managed by Terraform"

  # Add additional configuration as needed
}
