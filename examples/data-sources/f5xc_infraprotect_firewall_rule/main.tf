# Example configuration for f5xc_infraprotect_firewall_rule data source

data "f5xc_infraprotect_firewall_rule" "example" {
  name      = "existing-infraprotect_firewall_rule"
  namespace = "system"
}

output "infraprotect_firewall_rule_id" {
  value = data.f5xc_infraprotect_firewall_rule.example.id
}
