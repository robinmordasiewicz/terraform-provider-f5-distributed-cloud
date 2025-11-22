# Example configuration for f5xc_app_api_group

resource "f5xc_app_api_group" "example" {
  name        = "example-app_api_group"
  namespace   = "system"
  description = "Example AppAPIGroup resource managed by Terraform"

  # Add additional configuration as needed
}
