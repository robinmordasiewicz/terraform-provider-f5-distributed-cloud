# Example configuration for f5xc_cloud_connect

resource "f5xc_cloud_connect" "example" {
  name        = "example-cloud_connect"
  namespace   = "system"
  description = "Example CloudConnect resource managed by Terraform"

  # Add additional configuration as needed
}
