# Example configuration for f5_distributed_cloud_dns_lb_pool

resource "f5_distributed_cloud_dns_lb_pool" "example" {
  name        = "example-dns_lb_pool"
  namespace   = "system"
  description = "Example DNSLBPool resource managed by Terraform"

  # Add additional configuration as needed
}
