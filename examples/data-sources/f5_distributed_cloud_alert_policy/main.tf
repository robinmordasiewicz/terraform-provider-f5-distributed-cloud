# Example configuration for f5_distributed_cloud_alert_policy data source

data "f5_distributed_cloud_alert_policy" "example" {
  name      = "existing-alert_policy"
  namespace = "system"
}

output "alert_policy_id" {
  value = data.f5_distributed_cloud_alert_policy.example.id
}
