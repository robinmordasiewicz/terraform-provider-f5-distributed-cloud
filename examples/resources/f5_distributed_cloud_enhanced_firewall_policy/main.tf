# Example configuration for f5_distributed_cloud_enhanced_firewall_policy

resource "f5_distributed_cloud_enhanced_firewall_policy" "example" {
  name        = "example-enhanced_firewall_policy"
  namespace   = "system"
  description = "Example EnhancedFirewallPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
