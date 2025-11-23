# Example configuration for f5_distributed_cloud_dns_lb_pool data source

data "f5_distributed_cloud_dns_lb_pool" "example" {
  name      = "existing-dns_lb_pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.f5_distributed_cloud_dns_lb_pool.example.id
}
