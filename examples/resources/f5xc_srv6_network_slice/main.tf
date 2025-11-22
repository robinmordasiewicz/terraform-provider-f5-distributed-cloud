# Example configuration for f5xc_srv6_network_slice

resource "f5xc_srv6_network_slice" "example" {
  name        = "example-srv6_network_slice"
  namespace   = "system"
  description = "Example Srv6NetworkSlice resource managed by Terraform"

  # Add additional configuration as needed
}
