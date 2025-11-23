# Example configuration for f5_distributed_cloud_mobile_base_config

resource "f5_distributed_cloud_mobile_base_config" "example" {
  name        = "example-mobile_base_config"
  namespace   = "system"
  description = "Example MobileBaseConfig resource managed by Terraform"

  # Add additional configuration as needed
}
