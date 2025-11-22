# Example configuration for f5xc_waf_exclusion_policy data source

data "f5xc_waf_exclusion_policy" "example" {
  name      = "existing-waf_exclusion_policy"
  namespace = "system"
}

output "waf_exclusion_policy_id" {
  value = data.f5xc_waf_exclusion_policy.example.id
}
