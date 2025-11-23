# Example configuration for f5_distributed_cloud_api_testing

resource "f5_distributed_cloud_api_testing" "example" {
  name        = "example-api_testing"
  namespace   = "system"
  description = "Example APITesting resource managed by Terraform"

  # Add additional configuration as needed
}
