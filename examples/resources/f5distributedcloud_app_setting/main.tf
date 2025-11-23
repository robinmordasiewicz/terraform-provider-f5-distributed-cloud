# Example configuration for f5distributedcloud_app_setting

resource "f5distributedcloud_app_setting" "example" {
  name        = "example-app_setting"
  namespace   = "system"
  description = "Example AppSetting resource managed by Terraform"

  # Add additional configuration as needed
}
