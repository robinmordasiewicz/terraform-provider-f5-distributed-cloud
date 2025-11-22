# Example configuration for f5xc_dns_load_balancer data source

data "f5xc_dns_load_balancer" "example" {
  name      = "existing-dns_load_balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.f5xc_dns_load_balancer.example.id
}
