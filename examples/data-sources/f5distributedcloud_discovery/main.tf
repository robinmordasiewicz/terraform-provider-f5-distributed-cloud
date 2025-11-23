# Example configuration for f5distributedcloud_discovery data source

data "f5distributedcloud_discovery" "example" {
  name      = "existing-discovery"
  namespace = "system"
}

output "discovery_id" {
  value = data.f5distributedcloud_discovery.example.id
}
