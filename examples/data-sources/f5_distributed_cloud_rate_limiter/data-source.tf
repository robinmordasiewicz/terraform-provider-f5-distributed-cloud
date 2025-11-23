# Example: Rate Limiter Data Source
# This example retrieves information about an existing Rate Limiter

data "f5_distributed_cloud_rate_limiter" "example" {
  name      = "my-rate-limiter"
  namespace = "my-namespace"
}

# Output the rate limiter details
output "rate_limiter_id" {
  description = "ID of the Rate Limiter"
  value       = data.f5_distributed_cloud_rate_limiter.example.id
}

output "rate_limiter_description" {
  description = "Description of the Rate Limiter"
  value       = data.f5_distributed_cloud_rate_limiter.example.description
}

output "rate_limit" {
  description = "Rate limit configuration"
  value       = "${data.f5_distributed_cloud_rate_limiter.example.total_number} per ${data.f5_distributed_cloud_rate_limiter.example.unit}"
}

output "burst_multiplier" {
  description = "Burst capacity multiplier"
  value       = data.f5_distributed_cloud_rate_limiter.example.burst_multiplier
}

# Use the data source to reference an existing rate limiter
# Rate limiters are commonly applied to HTTP load balancers
resource "f5_distributed_cloud_http_loadbalancer" "protected" {
  name        = "rate-limited-app"
  namespace   = data.f5_distributed_cloud_rate_limiter.example.namespace
  description = "Load balancer protected by ${data.f5_distributed_cloud_rate_limiter.example.name}"

  domains = ["api.example.com"]

  http {
    port = 80
  }

  default_route_pools {
    pool {
      name      = "backend-pool"
      namespace = data.f5_distributed_cloud_rate_limiter.example.namespace
    }
    weight = 1
  }

  rate_limiter {
    name      = data.f5_distributed_cloud_rate_limiter.example.name
    namespace = data.f5_distributed_cloud_rate_limiter.example.namespace
  }
}
