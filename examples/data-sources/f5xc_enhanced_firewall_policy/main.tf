# Example configuration for f5xc_enhanced_firewall_policy data source

data "f5xc_enhanced_firewall_policy" "example" {
  name      = "existing-enhanced_firewall_policy"
  namespace = "system"
}

output "enhanced_firewall_policy_id" {
  value = data.f5xc_enhanced_firewall_policy.example.id
}
