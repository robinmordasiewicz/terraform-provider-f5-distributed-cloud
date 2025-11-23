# Example configuration for f5distributedcloud_token

resource "f5distributedcloud_token" "example" {
  name        = "example-token"
  namespace   = "system"
  description = "Example Token resource managed by Terraform"

  # Add additional configuration as needed
}
