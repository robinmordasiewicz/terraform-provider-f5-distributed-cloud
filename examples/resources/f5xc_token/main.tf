# Example configuration for f5xc_token

resource "f5xc_token" "example" {
  name        = "example-token"
  namespace   = "system"
  description = "Example Token resource managed by Terraform"

  # Add additional configuration as needed
}
