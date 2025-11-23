# Example configuration for f5_distributed_cloud_infraprotect_firewall_ruleset data source

data "f5_distributed_cloud_infraprotect_firewall_ruleset" "example" {
  name      = "existing-infraprotect_firewall_ruleset"
  namespace = "system"
}

output "infraprotect_firewall_ruleset_id" {
  value = data.f5_distributed_cloud_infraprotect_firewall_ruleset.example.id
}
