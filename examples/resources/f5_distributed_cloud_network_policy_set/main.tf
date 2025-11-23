# Example: Network Policy Set
resource "f5_distributed_cloud_network_policy_set" "example" {
  name        = "my-network-policy-set"
  namespace   = "my-namespace"
  description = "Example network policy set grouping policies"
}
