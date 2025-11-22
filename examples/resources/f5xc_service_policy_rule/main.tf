# Example configuration for f5xc_service_policy_rule

resource "f5xc_service_policy_rule" "example" {
  name        = "example-service_policy_rule"
  namespace   = "system"
  description = "Example ServicePolicyRule resource managed by Terraform"

  # Add additional configuration as needed
}
