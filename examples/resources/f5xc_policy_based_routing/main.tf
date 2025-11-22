# Example configuration for f5xc_policy_based_routing

resource "f5xc_policy_based_routing" "example" {
  name        = "example-policy_based_routing"
  namespace   = "system"
  description = "Example PolicyBasedRouting resource managed by Terraform"

  # Add additional configuration as needed
}
