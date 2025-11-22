# Example: API Group Element
resource "f5xc_api_group_element" "example" {
  name        = "my-api-group-element"
  namespace   = "my-namespace"
  description = "Example API group element for endpoint definition"
}
