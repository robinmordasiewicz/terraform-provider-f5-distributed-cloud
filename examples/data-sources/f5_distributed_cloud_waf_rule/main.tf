# Example: Look up an existing WAF rule
data "f5_distributed_cloud_waf_rule" "example" {
  name      = "my-waf-rule"
  namespace = "my-namespace"
}

# Output the WAF rule mode (BLOCK or DETECT)
output "waf_mode" {
  value = data.f5_distributed_cloud_waf_rule.example.mode
}

output "waf_rule_id" {
  value = data.f5_distributed_cloud_waf_rule.example.id
}
