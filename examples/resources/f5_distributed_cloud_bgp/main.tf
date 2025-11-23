# Example configuration for f5_distributed_cloud_bgp

resource "f5_distributed_cloud_bgp" "example" {
  name        = "example-bgp"
  namespace   = "system"
  description = "Example BGP resource managed by Terraform"

  # Add additional configuration as needed
}
