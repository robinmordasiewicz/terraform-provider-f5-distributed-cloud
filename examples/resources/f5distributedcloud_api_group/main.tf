# Example: API Group
resource "f5distributedcloud_api_group" "example" {
  name        = "my-api-group"
  namespace   = "my-namespace"
  description = "Example API group for endpoint organization"
}
