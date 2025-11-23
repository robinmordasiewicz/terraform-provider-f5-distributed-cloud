# Example configuration for f5distributedcloud_dns_lb_pool data source

data "f5distributedcloud_dns_lb_pool" "example" {
  name      = "existing-dns_lb_pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.f5distributedcloud_dns_lb_pool.example.id
}
