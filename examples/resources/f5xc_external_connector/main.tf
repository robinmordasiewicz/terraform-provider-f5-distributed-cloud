# Example: External Connector
resource "f5xc_external_connector" "example" {
  name        = "my-external-connector"
  namespace   = "my-namespace"
  description = "Example external connector for integration"
}
