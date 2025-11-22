# Example configuration for f5xc_discovery data source

data "f5xc_discovery" "example" {
  name      = "existing-discovery"
  namespace = "system"
}

output "discovery_id" {
  value = data.f5xc_discovery.example.id
}
