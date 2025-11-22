# Example: Authentication
resource "f5xc_authentication" "example" {
  name        = "my-authentication"
  namespace   = "my-namespace"
  description = "Example authentication configuration"
}
