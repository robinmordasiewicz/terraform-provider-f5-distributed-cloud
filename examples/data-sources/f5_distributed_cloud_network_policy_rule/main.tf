# Example configuration for f5_distributed_cloud_network_policy_rule data source

data "f5_distributed_cloud_network_policy_rule" "example" {
  name      = "existing-network_policy_rule"
  namespace = "system"
}

output "network_policy_rule_id" {
  value = data.f5_distributed_cloud_network_policy_rule.example.id
}
