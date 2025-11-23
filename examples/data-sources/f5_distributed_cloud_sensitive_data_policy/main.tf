# Example configuration for f5_distributed_cloud_sensitive_data_policy data source

data "f5_distributed_cloud_sensitive_data_policy" "example" {
  name      = "existing-sensitive_data_policy"
  namespace = "system"
}

output "sensitive_data_policy_id" {
  value = data.f5_distributed_cloud_sensitive_data_policy.example.id
}
