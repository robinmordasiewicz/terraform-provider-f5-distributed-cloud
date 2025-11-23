# Example configuration for f5_distributed_cloud_dns_zone

resource "f5_distributed_cloud_dns_zone" "example" {
  name        = "example-dns_zone"
  namespace   = "system"
  description = "Example DNSZone resource managed by Terraform"

  # Add additional configuration as needed
}
