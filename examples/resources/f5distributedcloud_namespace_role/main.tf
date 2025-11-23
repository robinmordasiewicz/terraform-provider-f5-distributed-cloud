# Example: Namespace Role
resource "f5distributedcloud_namespace_role" "example" {
  name        = "my-namespace-role"
  namespace   = "my-namespace"
  description = "Example namespace role for permissions"
}
