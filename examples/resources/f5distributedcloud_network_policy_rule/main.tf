# Example configuration for f5distributedcloud_network_policy_rule

resource "f5distributedcloud_network_policy_rule" "example" {
  name        = "example-network_policy_rule"
  namespace   = "system"
  description = "Example NetworkPolicyRule resource managed by Terraform"

  # Add additional configuration as needed
}
