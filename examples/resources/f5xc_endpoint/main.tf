# Example: Endpoint
resource "f5xc_endpoint" "example" {
  name        = "my-endpoint"
  namespace   = "my-namespace"
  description = "Example endpoint for service destination"
}
