# Example configuration for f5distributedcloud_endpoint

resource "f5distributedcloud_endpoint" "example" {
  name        = "example-endpoint"
  namespace   = "system"
  description = "Example Endpoint resource managed by Terraform"

  # Add additional configuration as needed
}
