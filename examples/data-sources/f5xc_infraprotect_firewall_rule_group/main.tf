# Example configuration for f5xc_infraprotect_firewall_rule_group data source

data "f5xc_infraprotect_firewall_rule_group" "example" {
  name      = "existing-infraprotect_firewall_rule_group"
  namespace = "system"
}

output "infraprotect_firewall_rule_group_id" {
  value = data.f5xc_infraprotect_firewall_rule_group.example.id
}
