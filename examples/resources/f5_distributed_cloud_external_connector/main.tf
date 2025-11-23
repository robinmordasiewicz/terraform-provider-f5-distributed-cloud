# Example configuration for f5_distributed_cloud_external_connector

resource "f5_distributed_cloud_external_connector" "example" {
  name        = "example-external_connector"
  namespace   = "system"
  description = "Example ExternalConnector resource managed by Terraform"

  # Add additional configuration as needed
}
