# Example configuration for f5xc_cloud_credentials

resource "f5xc_cloud_credentials" "example" {
  name        = "example-cloud_credentials"
  namespace   = "system"
  description = "Example CloudCredentials resource managed by Terraform"

  # Add additional configuration as needed
}
