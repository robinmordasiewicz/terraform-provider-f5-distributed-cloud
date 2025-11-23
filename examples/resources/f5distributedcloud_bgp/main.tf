# Example configuration for f5distributedcloud_bgp

resource "f5distributedcloud_bgp" "example" {
  name        = "example-bgp"
  namespace   = "system"
  description = "Example BGP resource managed by Terraform"

  # Add additional configuration as needed
}
