# Example: API Group
resource "f5_distributed_cloud_api_group" "example" {
  name        = "my-api-group"
  namespace   = "my-namespace"
  description = "Example API group for endpoint organization"
}
