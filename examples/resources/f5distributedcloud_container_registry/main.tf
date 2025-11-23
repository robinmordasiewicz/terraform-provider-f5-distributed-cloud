# Example configuration for f5distributedcloud_container_registry

resource "f5distributedcloud_container_registry" "example" {
  name        = "example-container_registry"
  namespace   = "system"
  description = "Example ContainerRegistry resource managed by Terraform"

  # Add additional configuration as needed
}
