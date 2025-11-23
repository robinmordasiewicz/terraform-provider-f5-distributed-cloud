# Example configuration for f5distributedcloud_app_type

resource "f5distributedcloud_app_type" "example" {
  name        = "example-app_type"
  namespace   = "system"
  description = "Example AppType resource managed by Terraform"

  # Add additional configuration as needed
}
