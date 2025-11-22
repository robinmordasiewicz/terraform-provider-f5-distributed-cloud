# Example configuration for f5xc_mobile_base_config

resource "f5xc_mobile_base_config" "example" {
  name        = "example-mobile_base_config"
  namespace   = "system"
  description = "Example MobileBaseConfig resource managed by Terraform"

  # Add additional configuration as needed
}
