# Example configuration for f5distributedcloud_dns_lb_health_check

resource "f5distributedcloud_dns_lb_health_check" "example" {
  name        = "example-dns_lb_health_check"
  namespace   = "system"
  description = "Example DNSLBHealthCheck resource managed by Terraform"

  # Add additional configuration as needed
}
