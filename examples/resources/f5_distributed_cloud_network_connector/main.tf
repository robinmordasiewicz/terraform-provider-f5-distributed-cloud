# Example configuration for f5_distributed_cloud_network_connector

resource "f5_distributed_cloud_network_connector" "example" {
  name        = "example-network_connector"
  namespace   = "system"
  description = "Example NetworkConnector resource managed by Terraform"

  # Add additional configuration as needed
}
