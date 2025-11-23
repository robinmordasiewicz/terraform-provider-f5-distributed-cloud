# Example configuration for f5distributedcloud_dns_zone

resource "f5distributedcloud_dns_zone" "example" {
  name        = "example-dns_zone"
  namespace   = "system"
  description = "Example DNSZone resource managed by Terraform"

  # Add additional configuration as needed
}
