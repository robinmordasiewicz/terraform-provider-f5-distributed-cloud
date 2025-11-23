# Example: RBAC Policy
resource "f5_distributed_cloud_rbac_policy" "example" {
  name        = "my-rbac-policy"
  namespace   = "my-namespace"
  description = "Example RBAC policy for access control"
}
