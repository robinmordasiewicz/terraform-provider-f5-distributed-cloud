# Example configuration for f5_distributed_cloud_bot_network_policy data source

data "f5_distributed_cloud_bot_network_policy" "example" {
  name      = "existing-bot_network_policy"
  namespace = "system"
}

output "bot_network_policy_id" {
  value = data.f5_distributed_cloud_bot_network_policy.example.id
}
