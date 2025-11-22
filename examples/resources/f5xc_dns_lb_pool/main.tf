# Example configuration for f5xc_dns_lb_pool

resource "f5xc_dns_lb_pool" "example" {
  name        = "example-dns_lb_pool"
  namespace   = "system"
  description = "Example DNSLBPool resource managed by Terraform"

  # Add additional configuration as needed
}
