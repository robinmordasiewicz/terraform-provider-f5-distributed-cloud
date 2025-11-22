# Example configuration for f5xc_infraprotect

resource "f5xc_infraprotect" "example" {
  name        = "example-infraprotect"
  namespace   = "system"
  description = "Example Infraprotect resource managed by Terraform"

  # Add additional configuration as needed
}
