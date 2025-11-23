# Example configuration for f5_distributed_cloud_origin_pool data source

data "f5_distributed_cloud_origin_pool" "example" {
  name      = "existing-origin_pool"
  namespace = "system"
}

output "origin_pool_id" {
  value = data.f5_distributed_cloud_origin_pool.example.id
}
