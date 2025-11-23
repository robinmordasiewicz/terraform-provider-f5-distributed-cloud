# Example configuration for f5_distributed_cloud_endpoint

resource "f5_distributed_cloud_endpoint" "example" {
  name        = "example-endpoint"
  namespace   = "system"
  description = "Example Endpoint resource managed by Terraform"

  # Add additional configuration as needed
}
