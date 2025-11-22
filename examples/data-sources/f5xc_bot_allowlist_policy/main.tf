# Example configuration for f5xc_bot_allowlist_policy data source

data "f5xc_bot_allowlist_policy" "example" {
  name      = "existing-bot_allowlist_policy"
  namespace = "system"
}

output "bot_allowlist_policy_id" {
  value = data.f5xc_bot_allowlist_policy.example.id
}
