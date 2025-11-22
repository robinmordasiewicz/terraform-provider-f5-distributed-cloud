# Example configuration for f5xc_sensitive_data_policy

resource "f5xc_sensitive_data_policy" "example" {
  name        = "example-sensitive_data_policy"
  namespace   = "system"
  description = "Example SensitiveDataPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
