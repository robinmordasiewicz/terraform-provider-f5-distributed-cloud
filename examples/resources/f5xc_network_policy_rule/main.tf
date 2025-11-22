# Example configuration for f5xc_network_policy_rule

resource "f5xc_network_policy_rule" "example" {
  name        = "example-network_policy_rule"
  namespace   = "system"
  description = "Example NetworkPolicyRule resource managed by Terraform"

  # Add additional configuration as needed
}
