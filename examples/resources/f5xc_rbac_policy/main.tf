# Example: RBAC Policy
resource "f5xc_rbac_policy" "example" {
  name        = "my-rbac-policy"
  namespace   = "my-namespace"
  description = "Example RBAC policy for access control"
}
