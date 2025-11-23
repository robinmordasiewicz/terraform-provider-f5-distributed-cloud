# Example configuration for f5distributedcloud_dns_load_balancer

resource "f5distributedcloud_dns_load_balancer" "example" {
  name        = "example-dns_load_balancer"
  namespace   = "system"
  description = "Example DNSLoadBalancer resource managed by Terraform"

  # Add additional configuration as needed
}
