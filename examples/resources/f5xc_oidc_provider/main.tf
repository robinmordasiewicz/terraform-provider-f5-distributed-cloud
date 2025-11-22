# Example: OIDC Provider
resource "f5xc_oidc_provider" "example" {
  name        = "my-oidc-provider"
  namespace   = "my-namespace"
  description = "Example OIDC provider for authentication"
}
