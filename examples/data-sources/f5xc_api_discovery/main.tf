# Example configuration for f5xc_api_discovery data source

data "f5xc_api_discovery" "example" {
  name      = "existing-api_discovery"
  namespace = "system"
}

output "api_discovery_id" {
  value = data.f5xc_api_discovery.example.id
}
