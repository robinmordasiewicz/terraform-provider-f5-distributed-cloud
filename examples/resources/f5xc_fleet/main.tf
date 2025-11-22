# Example configuration for f5xc_fleet

resource "f5xc_fleet" "example" {
  name        = "example-fleet"
  namespace   = "system"
  description = "Example Fleet resource managed by Terraform"

  # Add additional configuration as needed
}
