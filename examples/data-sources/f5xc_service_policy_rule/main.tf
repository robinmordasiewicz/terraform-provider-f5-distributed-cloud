# Example configuration for f5xc_service_policy_rule data source

data "f5xc_service_policy_rule" "example" {
  name      = "existing-service_policy_rule"
  namespace = "system"
}

output "service_policy_rule_id" {
  value = data.f5xc_service_policy_rule.example.id
}
