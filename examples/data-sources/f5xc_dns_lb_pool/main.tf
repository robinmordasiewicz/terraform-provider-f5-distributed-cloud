# Example: Look up an existing DNS LB Pool
data "f5xc_dns_lb_pool" "example" {
  name      = "my-dns-lb-pool"
  namespace = "my-namespace"
}

output "pool_type" {
  value = data.f5xc_dns_lb_pool.example.pool_type
}
