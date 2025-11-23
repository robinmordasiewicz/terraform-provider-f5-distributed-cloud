# Example configuration for f5distributedcloud_authentication

resource "f5distributedcloud_authentication" "example" {
  name        = "example-authentication"
  namespace   = "system"
  description = "Example Authentication resource managed by Terraform"

  # Add additional configuration as needed
}
