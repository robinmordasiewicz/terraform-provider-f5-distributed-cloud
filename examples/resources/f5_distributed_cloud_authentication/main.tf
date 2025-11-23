# Example configuration for f5_distributed_cloud_authentication

resource "f5_distributed_cloud_authentication" "example" {
  name        = "example-authentication"
  namespace   = "system"
  description = "Example Authentication resource managed by Terraform"

  # Add additional configuration as needed
}
