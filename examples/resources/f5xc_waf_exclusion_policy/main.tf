# Example: WAF Exclusion Policy
resource "f5xc_waf_exclusion_policy" "example" {
  name        = "my-waf-exclusion-policy"
  namespace   = "my-namespace"
  description = "Example WAF exclusion policy for bypassing specific rules"
}
