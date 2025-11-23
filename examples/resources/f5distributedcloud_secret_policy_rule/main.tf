# Example configuration for f5distributedcloud_secret_policy_rule

resource "f5distributedcloud_secret_policy_rule" "example" {
  name        = "example-secret_policy_rule"
  namespace   = "system"
  description = "Example SecretPolicyRule resource managed by Terraform"

  # Add additional configuration as needed
}
