# Example configuration for f5_distributed_cloud_bgp_routing_policy

resource "f5_distributed_cloud_bgp_routing_policy" "example" {
  name        = "example-bgp_routing_policy"
  namespace   = "system"
  description = "Example BGPRoutingPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
