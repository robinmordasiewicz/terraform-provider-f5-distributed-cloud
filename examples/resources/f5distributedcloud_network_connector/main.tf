# Example configuration for f5distributedcloud_network_connector

resource "f5distributedcloud_network_connector" "example" {
  name        = "example-network_connector"
  namespace   = "system"
  description = "Example NetworkConnector resource managed by Terraform"

  # Add additional configuration as needed
}
