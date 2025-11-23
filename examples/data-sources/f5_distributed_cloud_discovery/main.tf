# Example configuration for f5_distributed_cloud_discovery data source

data "f5_distributed_cloud_discovery" "example" {
  name      = "existing-discovery"
  namespace = "system"
}

output "discovery_id" {
  value = data.f5_distributed_cloud_discovery.example.id
}
