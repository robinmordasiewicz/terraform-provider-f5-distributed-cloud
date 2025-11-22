# Example: Look up an existing DNS zone
data "f5xc_dns_zone" "example" {
  name      = "my-dns-zone"
  namespace = "my-namespace"
}

# Output the domain name
output "domain" {
  value = data.f5xc_dns_zone.example.domain
}

output "dns_zone_id" {
  value = data.f5xc_dns_zone.example.id
}
