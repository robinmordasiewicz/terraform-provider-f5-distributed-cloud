# Example configuration for f5_distributed_cloud_dns_zone data source

data "f5_distributed_cloud_dns_zone" "example" {
  name      = "existing-dns_zone"
  namespace = "system"
}

output "dns_zone_id" {
  value = data.f5_distributed_cloud_dns_zone.example.id
}
