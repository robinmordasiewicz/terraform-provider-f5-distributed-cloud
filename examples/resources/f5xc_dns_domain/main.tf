# Example configuration for f5xc_dns_domain

resource "f5xc_dns_domain" "example" {
  name        = "example-dns_domain"
  namespace   = "system"
  description = "Example DNSDomain resource managed by Terraform"

  # Add additional configuration as needed
}
