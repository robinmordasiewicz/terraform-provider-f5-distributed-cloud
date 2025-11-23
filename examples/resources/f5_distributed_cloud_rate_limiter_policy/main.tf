# Example configuration for f5_distributed_cloud_rate_limiter_policy

resource "f5_distributed_cloud_rate_limiter_policy" "example" {
  name        = "example-rate_limiter_policy"
  namespace   = "system"
  description = "Example RateLimiterPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
