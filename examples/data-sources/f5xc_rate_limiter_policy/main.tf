# Example configuration for f5xc_rate_limiter_policy data source

data "f5xc_rate_limiter_policy" "example" {
  name      = "existing-rate_limiter_policy"
  namespace = "system"
}

output "rate_limiter_policy_id" {
  value = data.f5xc_rate_limiter_policy.example.id
}
