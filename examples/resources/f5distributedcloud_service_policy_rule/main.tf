# Example configuration for f5distributedcloud_service_policy_rule

resource "f5distributedcloud_service_policy_rule" "example" {
  name        = "example-service_policy_rule"
  namespace   = "system"
  description = "Example ServicePolicyRule resource managed by Terraform"

  # Add additional configuration as needed
}
