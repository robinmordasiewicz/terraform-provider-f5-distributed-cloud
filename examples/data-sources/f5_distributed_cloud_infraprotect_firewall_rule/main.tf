# Example configuration for f5_distributed_cloud_infraprotect_firewall_rule data source

data "f5_distributed_cloud_infraprotect_firewall_rule" "example" {
  name      = "existing-infraprotect_firewall_rule"
  namespace = "system"
}

output "infraprotect_firewall_rule_id" {
  value = data.f5_distributed_cloud_infraprotect_firewall_rule.example.id
}
