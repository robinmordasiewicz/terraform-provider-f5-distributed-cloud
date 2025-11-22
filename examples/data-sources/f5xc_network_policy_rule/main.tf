# Example configuration for f5xc_network_policy_rule data source

data "f5xc_network_policy_rule" "example" {
  name      = "existing-network_policy_rule"
  namespace = "system"
}

output "network_policy_rule_id" {
  value = data.f5xc_network_policy_rule.example.id
}
