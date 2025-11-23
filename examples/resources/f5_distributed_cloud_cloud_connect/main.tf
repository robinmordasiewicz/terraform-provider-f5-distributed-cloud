# Example configuration for f5_distributed_cloud_cloud_connect

resource "f5_distributed_cloud_cloud_connect" "example" {
  name        = "example-cloud_connect"
  namespace   = "system"
  description = "Example CloudConnect resource managed by Terraform"

  # Add additional configuration as needed
}
