# Example configuration for f5_distributed_cloud_rate_limiter data source

data "f5_distributed_cloud_rate_limiter" "example" {
  name      = "existing-rate_limiter"
  namespace = "system"
}

output "rate_limiter_id" {
  value = data.f5_distributed_cloud_rate_limiter.example.id
}
