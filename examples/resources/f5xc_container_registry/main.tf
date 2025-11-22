# Example: Container Registry
resource "f5xc_container_registry" "example" {
  name        = "my-container-registry"
  namespace   = "my-namespace"
  description = "Example container registry for image storage"
}
