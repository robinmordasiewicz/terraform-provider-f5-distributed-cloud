# Example configuration for f5_distributed_cloud_container_registry

resource "f5_distributed_cloud_container_registry" "example" {
  name        = "example-container_registry"
  namespace   = "system"
  description = "Example ContainerRegistry resource managed by Terraform"

  # Add additional configuration as needed
}
