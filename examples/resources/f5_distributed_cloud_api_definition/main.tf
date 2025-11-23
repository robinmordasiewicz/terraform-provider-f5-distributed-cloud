# Example configuration for f5_distributed_cloud_api_definition

resource "f5_distributed_cloud_api_definition" "example" {
  name        = "example-api_definition"
  namespace   = "system"
  description = "Example APIDefinition resource managed by Terraform"

  # Add additional configuration as needed
}
