# Example configuration for f5xc_bot_endpoint_policy data source

data "f5xc_bot_endpoint_policy" "example" {
  name      = "existing-bot_endpoint_policy"
  namespace = "system"
}

output "bot_endpoint_policy_id" {
  value = data.f5xc_bot_endpoint_policy.example.id
}
