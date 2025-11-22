# Example configuration for f5xc_quota

resource "f5xc_quota" "example" {
  name        = "example-quota"
  namespace   = "system"
  description = "Example Quota resource managed by Terraform"

  # Add additional configuration as needed
}
