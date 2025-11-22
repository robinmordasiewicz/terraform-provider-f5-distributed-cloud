resource "f5xc_dns_domain" "example" {
  name        = "my-dns-domain"
  namespace   = "system"
  description = "DNS domain configuration"
  domain      = "app.example.com"
}
