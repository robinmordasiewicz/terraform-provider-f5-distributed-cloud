# Example configuration for f5xc_fast_acl_rule data source

data "f5xc_fast_acl_rule" "example" {
  name      = "existing-fast_acl_rule"
  namespace = "system"
}

output "fast_acl_rule_id" {
  value = data.f5xc_fast_acl_rule.example.id
}
