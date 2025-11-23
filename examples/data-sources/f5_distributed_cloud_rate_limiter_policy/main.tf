# Example configuration for f5_distributed_cloud_rate_limiter_policy data source

data "f5_distributed_cloud_rate_limiter_policy" "example" {
  name      = "existing-rate_limiter_policy"
  namespace = "system"
}

output "rate_limiter_policy_id" {
  value = data.f5_distributed_cloud_rate_limiter_policy.example.id
}
