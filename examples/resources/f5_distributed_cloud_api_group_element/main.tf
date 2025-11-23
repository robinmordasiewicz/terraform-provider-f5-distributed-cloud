# Example: API Group Element
resource "f5_distributed_cloud_api_group_element" "example" {
  name        = "my-api-group-element"
  namespace   = "my-namespace"
  description = "Example API group element for endpoint definition"
}
