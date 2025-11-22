# Example: App Setting
resource "f5xc_app_setting" "example" {
  name        = "my-app-setting"
  namespace   = "my-namespace"
  description = "Example application setting configuration"
}
