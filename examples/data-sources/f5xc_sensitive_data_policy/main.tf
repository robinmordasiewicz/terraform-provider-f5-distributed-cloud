# Example configuration for f5xc_sensitive_data_policy data source

data "f5xc_sensitive_data_policy" "example" {
  name      = "existing-sensitive_data_policy"
  namespace = "system"
}

output "sensitive_data_policy_id" {
  value = data.f5xc_sensitive_data_policy.example.id
}
