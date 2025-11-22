# Example: Network Policy Rule
resource "f5xc_network_policy_rule" "example" {
  name        = "my-network-policy-rule"
  namespace   = "my-namespace"
  description = "Example network policy rule"
}
