# Example configuration for f5xc_network_policy data source

data "f5xc_network_policy" "example" {
  name      = "existing-network_policy"
  namespace = "system"
}

output "network_policy_id" {
  value = data.f5xc_network_policy.example.id
}
