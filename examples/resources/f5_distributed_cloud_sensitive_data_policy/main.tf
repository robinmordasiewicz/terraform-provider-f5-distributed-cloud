# Example configuration for f5_distributed_cloud_sensitive_data_policy

resource "f5_distributed_cloud_sensitive_data_policy" "example" {
  name        = "example-sensitive_data_policy"
  namespace   = "system"
  description = "Example SensitiveDataPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
