# Example configuration for f5distributedcloud_dns_zone data source

data "f5distributedcloud_dns_zone" "example" {
  name      = "existing-dns_zone"
  namespace = "system"
}

output "dns_zone_id" {
  value = data.f5distributedcloud_dns_zone.example.id
}
