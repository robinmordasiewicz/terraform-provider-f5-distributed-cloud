# Example: Namespace Role
resource "f5_distributed_cloud_namespace_role" "example" {
  name        = "my-namespace-role"
  namespace   = "my-namespace"
  description = "Example namespace role for permissions"
}
