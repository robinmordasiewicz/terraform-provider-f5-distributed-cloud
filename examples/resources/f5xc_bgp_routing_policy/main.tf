# Example: BGP Routing Policy
resource "f5xc_bgp_routing_policy" "example" {
  name        = "my-bgp-routing-policy"
  namespace   = "my-namespace"
  description = "Example BGP routing policy for route advertisement"
}
