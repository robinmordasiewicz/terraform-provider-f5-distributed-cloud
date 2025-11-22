# Example configuration for f5xc_origin_pool data source

data "f5xc_origin_pool" "example" {
  name      = "existing-origin_pool"
  namespace = "system"
}

output "origin_pool_id" {
  value = data.f5xc_origin_pool.example.id
}
