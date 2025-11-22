# Example configuration for f5xc_dns_lb_pool data source

data "f5xc_dns_lb_pool" "example" {
  name      = "existing-dns_lb_pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.f5xc_dns_lb_pool.example.id
}
