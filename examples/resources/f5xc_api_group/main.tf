# Example: API Group
resource "f5xc_api_group" "example" {
  name        = "my-api-group"
  namespace   = "my-namespace"
  description = "Example API group for endpoint organization"
}
