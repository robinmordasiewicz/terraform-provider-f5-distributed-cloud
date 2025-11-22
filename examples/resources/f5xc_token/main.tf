# Example: Token
resource "f5xc_token" "example" {
  name        = "my-token"
  namespace   = "my-namespace"
  description = "Example token for site registration"
}
