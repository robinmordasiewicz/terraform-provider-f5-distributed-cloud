# Example configuration for f5xc_secret_policy_rule

resource "f5xc_secret_policy_rule" "example" {
  name        = "example-secret_policy_rule"
  namespace   = "system"
  description = "Example SecretPolicyRule resource managed by Terraform"

  # Add additional configuration as needed
}
