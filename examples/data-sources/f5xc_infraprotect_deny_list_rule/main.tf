# Example configuration for f5xc_infraprotect_deny_list_rule data source

data "f5xc_infraprotect_deny_list_rule" "example" {
  name      = "existing-infraprotect_deny_list_rule"
  namespace = "system"
}

output "infraprotect_deny_list_rule_id" {
  value = data.f5xc_infraprotect_deny_list_rule.example.id
}
