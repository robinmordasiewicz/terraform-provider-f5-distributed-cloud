# Example configuration for f5_distributed_cloud_healthcheck

resource "f5_distributed_cloud_healthcheck" "example" {
  name        = "example-healthcheck"
  namespace   = "system"
  description = "Example Healthcheck resource managed by Terraform"

  # Add additional configuration as needed
}
