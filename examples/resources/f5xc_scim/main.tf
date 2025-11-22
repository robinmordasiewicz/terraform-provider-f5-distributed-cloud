# Example configuration for f5xc_scim

resource "f5xc_scim" "example" {
  name        = "example-scim"
  namespace   = "system"
  description = "Example Scim resource managed by Terraform"

  # Add additional configuration as needed
}
