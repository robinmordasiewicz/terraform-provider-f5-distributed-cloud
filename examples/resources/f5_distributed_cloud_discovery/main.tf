# Example configuration for f5_distributed_cloud_discovery

resource "f5_distributed_cloud_discovery" "example" {
  name        = "example-discovery"
  namespace   = "system"
  description = "Example Discovery resource managed by Terraform"

  # Add additional configuration as needed
}
