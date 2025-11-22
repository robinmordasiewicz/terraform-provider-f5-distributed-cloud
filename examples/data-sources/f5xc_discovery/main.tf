# Example: Look up an existing discovery configuration
data "f5xc_discovery" "example" {
  name      = "my-discovery"
  namespace = "my-namespace"
}

# Output the discovery type
output "discovery_type" {
  value = data.f5xc_discovery.example.discovery_type
}

output "discovery_id" {
  value = data.f5xc_discovery.example.id
}
