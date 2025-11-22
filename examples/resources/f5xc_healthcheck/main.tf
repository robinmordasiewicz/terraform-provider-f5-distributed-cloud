# Example configuration for f5xc_healthcheck

resource "f5xc_healthcheck" "example" {
  name        = "example-healthcheck"
  namespace   = "system"
  description = "Example Healthcheck resource managed by Terraform"

  # Add additional configuration as needed
}
