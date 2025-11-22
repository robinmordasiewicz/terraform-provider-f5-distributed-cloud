# Example: Rate Limiter Policy
resource "f5xc_rate_limiter_policy" "example" {
  name        = "my-rate-limiter-policy"
  namespace   = "my-namespace"
  description = "Example rate limiter policy for traffic control"
}
