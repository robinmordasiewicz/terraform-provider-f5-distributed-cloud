# Example configuration for f5distributedcloud_cloud_credentials

resource "f5distributedcloud_cloud_credentials" "example" {
  name        = "example-cloud_credentials"
  namespace   = "system"
  description = "Example CloudCredentials resource managed by Terraform"

  # Add additional configuration as needed
}
