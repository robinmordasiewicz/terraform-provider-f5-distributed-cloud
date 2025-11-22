# Example configuration for f5xc_alert_gen_policy data source

data "f5xc_alert_gen_policy" "example" {
  name      = "existing-alert_gen_policy"
  namespace = "system"
}

output "alert_gen_policy_id" {
  value = data.f5xc_alert_gen_policy.example.id
}
