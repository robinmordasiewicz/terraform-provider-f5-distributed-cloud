# Example configuration for f5xc_bgp

resource "f5xc_bgp" "example" {
  name        = "example-bgp"
  namespace   = "system"
  description = "Example BGP resource managed by Terraform"

  # Add additional configuration as needed
}
