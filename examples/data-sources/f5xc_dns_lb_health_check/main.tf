# Example configuration for f5xc_dns_lb_health_check data source

data "f5xc_dns_lb_health_check" "example" {
  name      = "existing-dns_lb_health_check"
  namespace = "system"
}

output "dns_lb_health_check_id" {
  value = data.f5xc_dns_lb_health_check.example.id
}
