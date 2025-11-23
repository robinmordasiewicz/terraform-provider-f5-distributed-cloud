# Example configuration for f5distributedcloud_origin_pool data source

data "f5distributedcloud_origin_pool" "example" {
  name      = "existing-origin_pool"
  namespace = "system"
}

output "origin_pool_id" {
  value = data.f5distributedcloud_origin_pool.example.id
}
