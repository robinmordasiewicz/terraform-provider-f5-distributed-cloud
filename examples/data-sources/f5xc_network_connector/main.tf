# Example configuration for f5xc_network_connector data source

data "f5xc_network_connector" "example" {
  name      = "existing-network_connector"
  namespace = "system"
}

output "network_connector_id" {
  value = data.f5xc_network_connector.example.id
}
