# Example configuration for f5xc_app_type

resource "f5xc_app_type" "example" {
  name        = "example-app_type"
  namespace   = "system"
  description = "Example AppType resource managed by Terraform"

  # Add additional configuration as needed
}
