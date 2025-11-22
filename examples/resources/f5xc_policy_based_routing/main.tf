# Example: Policy Based Routing
resource "f5xc_policy_based_routing" "example" {
  name        = "my-policy-based-routing"
  namespace   = "my-namespace"
  description = "Example policy based routing configuration"
}
