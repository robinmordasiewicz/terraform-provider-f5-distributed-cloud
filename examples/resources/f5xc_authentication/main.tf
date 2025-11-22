# Example configuration for f5xc_authentication

resource "f5xc_authentication" "example" {
  name        = "example-authentication"
  namespace   = "system"
  description = "Example Authentication resource managed by Terraform"

  # Add additional configuration as needed
}
