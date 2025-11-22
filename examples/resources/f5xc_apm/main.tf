# Example configuration for f5xc_apm

resource "f5xc_apm" "example" {
  name        = "example-apm"
  namespace   = "system"
  description = "Example Apm resource managed by Terraform"

  # Add additional configuration as needed
}
