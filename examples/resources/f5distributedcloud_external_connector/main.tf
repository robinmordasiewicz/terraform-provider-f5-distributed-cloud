# Example configuration for f5distributedcloud_external_connector

resource "f5distributedcloud_external_connector" "example" {
  name        = "example-external_connector"
  namespace   = "system"
  description = "Example ExternalConnector resource managed by Terraform"

  # Add additional configuration as needed
}
