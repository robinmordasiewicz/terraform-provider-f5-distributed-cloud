# Example configuration for f5distributedcloud_enhanced_firewall_policy

resource "f5distributedcloud_enhanced_firewall_policy" "example" {
  name        = "example-enhanced_firewall_policy"
  namespace   = "system"
  description = "Example EnhancedFirewallPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
