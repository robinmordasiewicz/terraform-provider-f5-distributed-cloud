# Example: Look up an existing enhanced firewall policy
data "f5xc_enhanced_firewall_policy" "example" {
  name      = "my-enhanced-firewall-policy"
  namespace = "my-namespace"
}

# Output the enabled status
output "enabled" {
  value = data.f5xc_enhanced_firewall_policy.example.enabled
}

output "enhanced_firewall_policy_id" {
  value = data.f5xc_enhanced_firewall_policy.example.id
}
