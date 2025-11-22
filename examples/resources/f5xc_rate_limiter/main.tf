# Example configuration for f5xc_rate_limiter

resource "f5xc_rate_limiter" "example" {
  name        = "example-rate_limiter"
  namespace   = "system"
  description = "Example RateLimiter resource managed by Terraform"

  # Add additional configuration as needed
}
