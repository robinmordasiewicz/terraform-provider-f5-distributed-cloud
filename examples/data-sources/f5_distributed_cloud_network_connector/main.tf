# Example configuration for f5_distributed_cloud_network_connector data source

data "f5_distributed_cloud_network_connector" "example" {
  name      = "existing-network_connector"
  namespace = "system"
}

output "network_connector_id" {
  value = data.f5_distributed_cloud_network_connector.example.id
}
