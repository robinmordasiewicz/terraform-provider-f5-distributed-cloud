# Example configuration for f5xc_network_policy_view

resource "f5xc_network_policy_view" "example" {
  name        = "example-network_policy_view"
  namespace   = "system"
  description = "Example NetworkPolicyView resource managed by Terraform"

  # Add additional configuration as needed
}
