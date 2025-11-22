# Example configuration for f5xc_bot_network_policy data source

data "f5xc_bot_network_policy" "example" {
  name      = "existing-bot_network_policy"
  namespace = "system"
}

output "bot_network_policy_id" {
  value = data.f5xc_bot_network_policy.example.id
}
