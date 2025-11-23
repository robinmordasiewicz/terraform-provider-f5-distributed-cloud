# Example configuration for f5_distributed_cloud_token

resource "f5_distributed_cloud_token" "example" {
  name        = "example-token"
  namespace   = "system"
  description = "Example Token resource managed by Terraform"

  # Add additional configuration as needed
}
