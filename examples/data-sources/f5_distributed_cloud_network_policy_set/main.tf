# Example configuration for f5_distributed_cloud_network_policy_set data source

data "f5_distributed_cloud_network_policy_set" "example" {
  name      = "existing-network_policy_set"
  namespace = "system"
}

output "network_policy_set_id" {
  value = data.f5_distributed_cloud_network_policy_set.example.id
}
