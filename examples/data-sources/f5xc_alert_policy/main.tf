# Example: Look up an existing alert policy
data "f5xc_alert_policy" "example" {
  name      = "my-alert-policy"
  namespace = "my-namespace"
}

# Output whether the alert policy is enabled
output "enabled" {
  value = data.f5xc_alert_policy.example.enabled
}

output "alert_policy_id" {
  value = data.f5xc_alert_policy.example.id
}
