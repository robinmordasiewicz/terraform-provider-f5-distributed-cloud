# Example configuration for f5distributedcloud_network_policy_set data source

data "f5distributedcloud_network_policy_set" "example" {
  name      = "existing-network_policy_set"
  namespace = "system"
}

output "network_policy_set_id" {
  value = data.f5distributedcloud_network_policy_set.example.id
}
