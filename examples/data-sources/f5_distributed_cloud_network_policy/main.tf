# Example configuration for f5_distributed_cloud_network_policy data source

data "f5_distributed_cloud_network_policy" "example" {
  name      = "existing-network_policy"
  namespace = "system"
}

output "network_policy_id" {
  value = data.f5_distributed_cloud_network_policy.example.id
}
