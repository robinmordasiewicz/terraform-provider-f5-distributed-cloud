# Example configuration for f5distributedcloud_dns_domain

resource "f5distributedcloud_dns_domain" "example" {
  name        = "example-dns_domain"
  namespace   = "system"
  description = "Example DNSDomain resource managed by Terraform"

  # Add additional configuration as needed
}
