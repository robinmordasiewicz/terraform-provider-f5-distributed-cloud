# Example configuration for f5_distributed_cloud_app_type

resource "f5_distributed_cloud_app_type" "example" {
  name        = "example-app_type"
  namespace   = "system"
  description = "Example AppType resource managed by Terraform"

  # Add additional configuration as needed
}
