# Example configuration for f5distributedcloud_bgp_routing_policy

resource "f5distributedcloud_bgp_routing_policy" "example" {
  name        = "example-bgp_routing_policy"
  namespace   = "system"
  description = "Example BGPRoutingPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
