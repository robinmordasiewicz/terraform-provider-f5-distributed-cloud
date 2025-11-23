# Example configuration for f5distributedcloud_healthcheck

resource "f5distributedcloud_healthcheck" "example" {
  name        = "example-healthcheck"
  namespace   = "system"
  description = "Example Healthcheck resource managed by Terraform"

  # Add additional configuration as needed
}
