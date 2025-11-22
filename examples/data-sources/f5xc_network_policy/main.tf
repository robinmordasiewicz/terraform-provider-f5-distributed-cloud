# Example: Look up an existing network policy
data "f5xc_network_policy" "example" {
  name      = "my-network-policy"
  namespace = "my-namespace"
}

# Output the network policy type
output "policy_type" {
  value = data.f5xc_network_policy.example.policy_type
}

output "network_policy_id" {
  value = data.f5xc_network_policy.example.id
}
