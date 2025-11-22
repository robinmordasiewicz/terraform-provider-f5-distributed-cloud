resource "f5xc_dns_zone" "example" {
  name        = "my-dns-zone"
  namespace   = "system"
  description = "Primary DNS zone for example.com"
  domain      = "example.com"
}
