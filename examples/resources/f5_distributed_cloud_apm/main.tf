# Example configuration for f5_distributed_cloud_apm

resource "f5_distributed_cloud_apm" "example" {
  name        = "example-apm"
  namespace   = "system"
  description = "Example Apm resource managed by Terraform"

  # Add additional configuration as needed
}
