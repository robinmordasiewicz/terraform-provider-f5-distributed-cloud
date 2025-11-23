# Example configuration for f5_distributed_cloud_secret_policy_rule data source

data "f5_distributed_cloud_secret_policy_rule" "example" {
  name      = "existing-secret_policy_rule"
  namespace = "system"
}

output "secret_policy_rule_id" {
  value = data.f5_distributed_cloud_secret_policy_rule.example.id
}
