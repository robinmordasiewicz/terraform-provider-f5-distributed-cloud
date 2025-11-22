# Example configuration for f5xc_v1_http_monitor

resource "f5xc_v1_http_monitor" "example" {
  name        = "example-v1_http_monitor"
  namespace   = "system"
  description = "Example V1HTTPMonitor resource managed by Terraform"

  # Add additional configuration as needed
}
